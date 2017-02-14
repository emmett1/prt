package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
)

// Diff lists outdated packages.
func diff(args []string) {
	// Define valid arguments.
	o := optparse.New()
	argn := o.Bool("no-alias", 'n', false)
	argv := o.Bool("version", 'v', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt diff [arguments]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -n,   --no-alias        disable aliasing")
		fmt.Println("  -v,   --version         print with version info")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	// Get all ports.
	all, err := portAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed ports.
	inst, err := portInst()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed port versions.
	instv, err := portInstVers()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i, p := range inst {
		// Get port location.
		ll, err := portLoc(all, p)
		if err != nil {
			continue
		}
		l := ll[0]

		// Alias if needed.
		if *argn {
			l = portAlias(l)
		}

		// Get available version and release from Pkgfile.
		if err := initPkgfile(portFullLoc(l), []string{"Version", "Release"}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		// Combine version and release.
		availv := pkgfile.Version + "-" + pkgfile.Release

		// Print if installed and available version don't match.
		if availv != instv[i] {
			fmt.Print(p)

			// Print version information if needed.
			if *argv {
				fmt.Print(" " + instv[i])

				color.Set(config.DarkColor)
				fmt.Print(" -> ")
				color.Unset()

				fmt.Print(availv)
			}

			fmt.Println()
		}
	}
}