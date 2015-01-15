package main

import (
	"bytes"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"gopkg.in/alecthomas/kingpin.v1"
	"os"
	"path/filepath"
)

var (
	verbose     = kingpin.Flag("verbose", "Enable verbose output.").Short('v').Bool()
	veryVerbose = kingpin.Flag("veryverbose", "Enable very verbose output").Short('V').Bool()
	pkgRoot     = kingpin.Flag("pkgroot", "The package path root. Defaults to $GOPATH. Only needs to be specified if $GOPATH is not a single path").Short('r').OverrideDefaultFromEnvar("GOPATH").ExistingDir()
	pkglist     = kingpin.Arg("pkgs", "A list of package paths in which to find \".dodl\" files").Strings()
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

type Generator struct {
	file     *os.File
	r        *reader
	buf      bytes.Buffer
	typedefs map[typeId]*typedef
}

func generate(pkg string) error {
	if *verbose || *veryVerbose {
		fmt.Printf("Scanning package %s\n", pkg)
	}
	pkgUUID, err := uuid.NewV5(uuid.NamespaceURL, []byte(pkg))
	if err != nil {
		return fmt.Errorf("Couldn't create package UUID for package %s, err: %s", pkg, err)
	}
	g := Generator{typedefs: make(map[typeId]*typedef, 100)}
	for _, fn := range listSourceFilenames(filepath.Join(*pkgRoot, "src", pkg)) {
		if *verbose || *veryVerbose {
			fmt.Printf("  Parsing %s - UUID:%v\n", fn, pkgUUID)
		}
		if err := g.parseFile(pkgUUID, fn); err != nil {
			return err
		}
	}
	if err := g.validate(); err != nil {
		return err
	}
	return nil
}

func listSourceFilenames(pkgDir string) []string {
	pattern := filepath.Join(pkgDir, "*.dodl")
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

func (g *Generator) parseFile(pkgUUID *uuid.UUID, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // f is readOnly, so no need to check for close errors.
	g.file = f
	g.r = newReader(f)
	for {
		var t *typedef
		t, err = g.parseTypedef(pkgUUID)
		if err != nil {
			return err
		}
		if t == nil {
			return nil
		}
		if old, present := g.typedefs[t.typeId]; present {
			return fmt.Errorf("Duplicate definition of type id %s on line %d of file %s\n   It is also defined on line %d of file %s", t.typeId.String(), t.srcline, t.srcfile, old.srcline, old.srcfile)
		}
		g.typedefs[t.typeId] = t
	}
}

func (g *Generator) validate() error {
	if err := g.validateTypenames(); err != nil {
		return err
	}
	return nil
}

func (g *Generator) validateTypenames() error {
	typenames := make(map[string]*typedef)
	for _, t := range g.typedefs {
		if old, present := typenames[t.name]; present {
			return fmt.Errorf("Duplicate definition of type name %s on line %d of file %s\n   It is also defined on line %d of file %s", t.name, t.srcline, t.srcfile, old.srcline, old.srcfile)
		}
		typenames[t.name] = t
	}
	return nil
}
