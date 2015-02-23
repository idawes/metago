// Copyright 2015 Ian Dawes. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// metago is a tool to .....
//
//
//

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/nu7hatch/gouuid"
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	verbose     = kingpin.Flag("verbose", "Enable verbose output.").Short('v').Bool()
	veryVerbose = kingpin.Flag("veryverbose", "Enable very verbose output").Short('V').Bool()
	pkgRoot     = kingpin.Flag("pkgroot", "The package path root. Defaults to $GOPATH. Only needs to be specified if $GOPATH is not a single path").Short('r').OverrideDefaultFromEnvar("GOPATH").ExistingDir()
	pkglist     = kingpin.Arg("pkgs", "A list of package paths in which to find \".mtgo\" files").Strings()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	for _, pkg := range *pkglist {
		if err := generate(pkg); err != nil {
			fmt.Println("\nERROR: ", err)
			os.Exit(-1)
		}
	}
}

// Note: We're using deferred error handling to simplify the main generation code flow. The first error that occurs during generation is
// captured and will prevent further steps from performing any real work.
type generator struct {
	pkg            string
	pkgUUID        *uuid.UUID
	typedefs       map[typeID]*typedef
	typedefsByName map[string]*typedef
	err            error
}

func generate(pkg string) error {
	if *verbose || *veryVerbose {
		fmt.Printf("Scanning package %s\n", pkg)
	}
	u, err := uuid.NewV5(uuid.NamespaceURL, []byte(pkg))
	if err != nil {
		return fmt.Errorf("couldn't create package UUID for package %s, err: %s", pkg, err)
	}
	g := generator{pkg: pkg, pkgUUID: u, typedefs: make(map[typeID]*typedef, 0)}
	g.parseFiles()
	g.validateTypes()
	g.generate()
	return g.err
}

func (g *generator) parseFiles() {
	for _, f := range listSourceFilenames(filepath.Join(*pkgRoot, "src", g.pkg)) {
		g.parseFile(f)
	}
}

func listSourceFilenames(pkgDir string) []string {
	pattern := filepath.Join(pkgDir, "*.mtgo")
	if *veryVerbose {
		fmt.Printf("  Executing match using pattern %s\n", pattern)
	}
	fileNames, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("    Pattern match failed, error: %s\n", err)
		return []string{}
	}
	if fileNames == nil {
		if *verbose || *veryVerbose {
			fmt.Printf("No source files found in %s\n", pkgDir)
		}
		return []string{}
	}
	return fileNames
}

func (g *generator) parseFile(f string) {
	if g.err != nil {
		return
	}
	if *verbose || *veryVerbose {
		fmt.Printf("  Parsing %s - UUID:%v\n", f, g.pkgUUID)
	}
	r := newReader(f)
	defer r.close()
	for {
		t, err := parseTypedef(g.pkgUUID, r)
		if err != nil {
			if err == io.EOF {
				return
			}
			g.err = err
			return
		}
		if old, present := g.typedefs[t.typeID]; present {
			g.err = fmt.Errorf("duplicate definition of type id %s on line %d of file %s\n   It is also defined on line %d of file %s", t.typeID.String(), t.srcline, t.srcfile, old.srcline, old.srcfile)
			return
		}
		g.typedefs[t.typeID] = t
	}
}

func (g *generator) validateTypes() {
	g.validateTypenames()
	g.validateTypeHierarchy()
}

func (g *generator) validateTypenames() {
	if g.err != nil {
		return
	}
	g.typedefsByName = make(map[string]*typedef)
	for _, t := range g.typedefs {
		if old, present := g.typedefsByName[t.name]; present {
			g.err = fmt.Errorf("duplicate definition of type name %s on line %d of file %s\n   It is also defined on line %d of file %s", t.name, t.srcline, t.srcfile, old.srcline, old.srcfile)
			return
		}
		g.typedefsByName[t.name] = t
	}
}

func (g *generator) validateTypeHierarchy() {
	for _, t := range g.typedefs {
		if g.err != nil {
			return
		}
		g.err = t.validateTypeHierarchy(g.typedefsByName)
	}
}

type typedefList []*typedef

func (l typedefList) Len() int           { return len(l) }
func (l typedefList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l typedefList) Less(i, j int) bool { return l[i].typeID.Compare(&l[j].typeID) < 0 }

const schemaHeader = `

import ( 
	"fmt"

	"github.com/idawes/metago"
	"github.com/nu7hatch/gouuid"
)

var MetagoPackageUUID uuid.UUID

func init() {
	id, err := uuid.ParseHex("%s")
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse package UUID for package %s"))
	}
	MetagoPackageUUID = *id
}

var (
`

func (g *generator) generate() {
	sw := newWriter(filepath.Join(*pkgRoot, "src", g.pkg), "schema")
	sw.printf(schemaHeader, g.pkgUUID.String(), g.pkg)
	sortedTypedefs := make(typedefList, 0)
	for _, t := range g.typedefs {
		sortedTypedefs = append(sortedTypedefs, t)
	}
	sort.Sort(sortedTypedefs)
	for _, t := range sortedTypedefs {
		if g.err != nil {
			return
		}
		w := newWriter(filepath.Join(*pkgRoot, "src", g.pkg), t.name)
		t.generate(w)
		t.generateSchema(sw)
		g.err = w.close()
	}
	sw.printf(")")
	if g.err == nil {
		g.err = sw.close()
	}
}
