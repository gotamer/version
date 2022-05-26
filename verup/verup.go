package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gotamer/version"
)

const usage = `
Up to date version information from git.
 - by create a file (version.go) at current dir.
 - or output to stdout to view or pipe.

Usage of verup:
`

var (
	save   = flag.Bool("s", false, "Save to File")
	toFile = flag.String("f", "version.go", "Save to File name")
	help   = flag.Bool("h", false, "Display help")
	vers   = flag.Bool("v", false, "Display version")
)

func main() {

	flag.Parse()

	if *help {
		fmt.Println(usage)
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *vers {
		fmt.Printf("Tag Version: %s\n", VerTag)
		fmt.Printf("Git Version: %s\n", VerGit)
		os.Exit(0)
	}

	var err error
	if *save {
		err = version.Save(*toFile)
	} else {
		err = version.Out()
	}
	if err != nil {
		log.Println("ERR Version : ", err.Error())
	}
}
