package main

import (
	"bytes"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"gopkg.in/alecthomas/kingpin.v1"
	"io"
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

type generator struct {
	file           *os.File
	r              *reader
	buf            bytes.Buffer
	err            error
	typedefs       map[typeID]*typedef
	typedefsByName map[string]*typedef
}

func generate(pkg string) error {
	if *verbose || *veryVerbose {
		fmt.Printf("Scanning package %s\n", pkg)
	}
	g := generator{typedefs: make(map[typeID]*typedef, 100)}
	if err := g.parseFiles(pkg); err != nil {
		return err
	}
	if err := g.validate(); err != nil {
		return err
	}
	if err := g.generate(pkg); err != nil {
		return err
	}
	return nil
}

func (g *generator) parseFiles(pkg string) error {
	pkgUUID, err := uuid.NewV5(uuid.NamespaceURL, []byte(pkg))
	if err != nil {
		return fmt.Errorf("Couldn't create package UUID for package %s, err: %s", pkg, err)
	}
	for _, fn := range listSourceFilenames(filepath.Join(*pkgRoot, "src", pkg)) {
		if *verbose || *veryVerbose {
			fmt.Printf("  Parsing %s - UUID:%v\n", fn, pkgUUID)
		}
		if err := g.parseFile(pkgUUID, fn); err != nil {
			return err
		}
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

func (g *generator) parseFile(pkgUUID *uuid.UUID, filename string) error {
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
		if old, present := g.typedefs[t.typeID]; present {
			return fmt.Errorf("Duplicate definition of type id %s on line %d of file %s\n   It is also defined on line %d of file %s", t.typeID.String(), t.srcline, t.srcfile, old.srcline, old.srcfile)
		}
		g.typedefs[t.typeID] = t
	}
}

func (g *generator) validate() error {
	if err := g.validateTypenames(); err != nil {
		return err
	}
	if err := g.validateTypeHierarchy(); err != nil {
		return nil
	}
	return nil
}

func (g *generator) validateTypenames() error {
	g.typedefsByName = make(map[string]*typedef)
	for _, t := range g.typedefs {
		if old, present := g.typedefsByName[t.name]; present {
			return fmt.Errorf("Duplicate definition of type name %s on line %d of file %s\n   It is also defined on line %d of file %s", t.name, t.srcline, t.srcfile, old.srcline, old.srcfile)
		}
		g.typedefsByName[t.name] = t
	}
	return nil
}

func (g *generator) validateTypeHierarchy() error {
	for _, t := range g.typedefs {
		if err := t.validateTypeHierarchy(g.typedefsByName); err != nil {
			return err
		}
	}
	return nil
}

func (g *generator) generate(pkg string) error {
	for _, t := range g.typedefs {
		if err := g.generateType(t, pkg); err != nil {
			return err
		}
	}
	return nil
}

func (g *generator) generateType(t *typedef, pkg string) (err error) {
	g.buf = bytes.Buffer{}
	g.err = nil
	f, err := os.Create(filepath.Join(*pkgRoot, "src", pkg, fmt.Sprintf("%s.go.tmp", t.name)))
	if err != nil {
		return err
	}
	t.generate(g)
	if err := g.writeTo(f); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func (g *generator) writeTo(w io.Writer) error {
	if g.err == nil {
		_, g.err = g.buf.WriteTo(w)
	}
	return g.err
}

func (g *generator) printf(format string, args ...interface{}) {
	if g.err == nil {
		fmt.Fprintf(&g.buf, format, args...)
	}
}
