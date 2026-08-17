// Command create_mod builds a Go module zip in proxy.golang.org layout.
//
// Usage: create_mod <module-path> <version> <source-dir> <output-zip>
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: create_mod <module-path> <version> <source-dir> <output-zip>")
		os.Exit(1)
	}
	modPath, version, srcDir, outZip := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	if err := os.MkdirAll(filepath.Dir(outZip), 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "mkdir:", err)
		os.Exit(1)
	}

	f, err := os.Create(outZip)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create:", err)
		os.Exit(1)
	}
	defer f.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, srcDir); err != nil {
		fmt.Fprintln(os.Stderr, "CreateFromDir:", err)
		os.Exit(1)
	}
	fmt.Println("wrote", outZip)
}
