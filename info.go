package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go2c/optparse"
)

// info prints port information.
func info(args []string) {
	// Enable all arguments if the user hasn't specified any.
	var b bool
	if len(args) == 0 {
		b = true
	}

	// Define valid arguments.
	o := optparse.New()
	argd := o.Bool("description", 'd', b)
	argu := o.Bool("url", 'u', b)
	argm := o.Bool("maintainer", 'm', b)
	arge := o.Bool("depends", 'e', b)
	argo := o.Bool("optional", 'o', b)
	argv := o.Bool("version", 'v', b)
	argr := o.Bool("release", 'r', b)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt info [arguments]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -d,   --description     print description")
		fmt.Println("  -u,   --url             print url")
		fmt.Println("  -m,   --maintainer      print maintainer")
		fmt.Println("  -e,   --depends         print dependencies")
		fmt.Println("  -o,   --optional        print optional dependencies")
		fmt.Println("  -v,   --version         print version")
		fmt.Println("  -r,   --release         print release")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	if err := initPkgfile(".", []string{"Description", "URL", "Maintainer", "Depends", "Optional", "Version", "Release"}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Print info from Pkgfile.
	if *argd {
		fmt.Println("Description: " + pkgfile.Description)
	}
	if *argu {
		fmt.Println("URL: " + pkgfile.URL)
	}
	if *argm {
		fmt.Println("Maintainer: " + pkgfile.Maintainer)
	}
	if *arge {
		fmt.Println("Depends on: " + strings.Join(pkgfile.Depends, ", "))
	}
	if *argo {
		fmt.Println("Nice to have: " + strings.Join(pkgfile.Optional, ", "))
	}
	if *argv {
		fmt.Println("Version: " + pkgfile.Version)
	}
	if *argr {
		fmt.Println("Release: " + pkgfile.Release)
	}
}