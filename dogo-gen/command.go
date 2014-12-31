package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v1"
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
		pkgDir := filepath.Join(*pkgRoot, "src", pkg)
		if *verbose || *veryVerbose {
			fmt.Printf("Scanning package %s\n", pkgDir)
		}
		generate(pkgDir)
	}
}

func generate(pkgDir string) {
	for _, srcFilename := range listDODLFiles(pkgDir) {
		if *verbose || *veryVerbose {
			fmt.Printf("  Parsing %s\n", srcFilename)
		}
	}
}

func listDODLFiles(pkgDir string) []string {
	pattern := filepath.Join(pkgDir, "*.dodl")
	if *verbose || *veryVerbose {
		fmt.Printf("  Executing match using pattern %s\n", pattern)
	}
	fileNames, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("    Pattern match failed, error: %s\n", err)
		return []string{}
	}
	if fileNames == nil {
		if *verbose || *veryVerbose {
			fmt.Printf("No .mdo files found in %s\n", pkgDir)
		}
		return []string{}
	}
	return fileNames
}
