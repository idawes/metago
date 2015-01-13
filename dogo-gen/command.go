package main

import (
	"bytes"
	"fmt"
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
	var g Generator
	for _, pkg := range *pkglist {
		g.generate(filepath.Join(*pkgRoot, "src", pkg))
	}
}

type Generator struct {
	file     *os.File
	r        *reader
	buf      bytes.Buffer
	typedefs []*typedef
}

func (g *Generator) generate(pkgDir string) {
	if *verbose || *veryVerbose {
		fmt.Printf("Scanning package %s\n", pkgDir)
	}
	g.buf.Reset()
	g.typedefs = make([]*typedef, 100)
	for _, fn := range listSourceFileNames(pkgDir) {
		if *verbose || *veryVerbose {
			fmt.Printf("  Parsing %s\n", fn)
		}
		if err := g.parseFile(fn); err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

func listSourceFileNames(pkgDir string) []string {
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

func (g *Generator) parseFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // f is readOnly, so no need to check for close errors.
	g.file = f
	g.r = newReader(f)
	for {
		var t *typedef
		t, err = g.parseTypedef()
		if err != nil {
			return err
		}
		if t == nil {
			return nil
		}
		g.typedefs = append(g.typedefs, t)
	}
	return nil
}
