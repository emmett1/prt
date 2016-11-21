package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/util"
)

// Initialize global variables
var AllPorts, Checked, InstPorts []string
var Iteration int

// This function prints dependencies recursivly
func recursive(path string, a, t bool) {
	// Read out Pkgfile
	pkgfile, err := ioutil.ReadFile(path + "/Pkgfile")
	if err != nil {
		fmt.Println("Could not read Pkgfile.")
		os.Exit(1)
	}

	// Read out Pkgfile dependencies
	deps := util.ReadDepends(pkgfile, "Depends on")

	var locList []string
	var loc string

	for _, dep := range deps {
		// Continue if already checked
		if util.StringInList(dep, Checked) {
			continue
		} else {
			Checked = append(Checked, dep)
		}

		// Continue if already installed
		if !a {
			if util.StringInList(dep, InstPorts) {
				continue
			}
		}

		// Get port location
		locList = util.GetPortLoc(AllPorts, dep)
		if len(locList) < 1 {
			return
		} else {
			loc = locList[0]
		}

		// Print tree arrowsPrt
		if t {
			if Iteration > 0 {
				color.Set(color.FgBlack, color.Bold)
				fmt.Printf(strings.Repeat("-  ", Iteration))
				color.Unset()
			}
			Iteration += 1
		}

		// Finally print the port :)
		fmt.Println(loc)

		// Loop
		recursive(config.Struct.PrtDir+"/"+loc, a, t)

		if t {
			Iteration -= 1
		}
	}
}

func Depends(args []string) {
	// Define opts
	shortopts := "hadt"
	longopts := []string{
		"--help",
		"--disable-alias",
		"--tree",
	}

	// Read out opts
	opts, _, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Println("Invaild argument, use -h for a list of arguments.")
		os.Exit(1)
	}

	// Initialize opt variables
	//var a, d, t bool
	var a, t bool

	for _, opt := range opts {
		switch opt[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt depends [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -a,   --all             also list installed dependencies")
			fmt.Println("  -d,   --disable-alias   disable aliasing")
			fmt.Println("  -t,   --tree            list using tree view")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-a", "--all":
			a = true
		case "-d", "--disable-alias":
			//d = true
		case "-t", "--tree":
			t = true
		}
	}

	AllPorts = util.ListAllPorts()
	if !a {
		InstPorts = util.ListInstPorts()
	}

	recursive("./", a, t)
}
