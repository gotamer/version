package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var version struct {
	Long  string
	Tags  string
	Git   string
	Dirty bool
}

func main() {
	cmd := exec.Command("git", "describe", "--always", "--long", "--tags", "--dirty")
	cmd.Stdin = strings.NewReader("Version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	version.Long = out.String()
	vers := strings.Split(version.Long, "-")
	l := len(vers)
	if strings.Contains(vers[l-1], "dirty") {
		fmt.Printf("Is Dirty")
		version.Dirty = true
		vers = vers[0 : l-1]
		fmt.Printf("Contants dirty new vers %s", vers)
	}
	l = len(vers)
	version.Git = vers[l-1]
	vers = vers[0 : l-1]
	l = len(vers)

	if len(vers) == 2 {
		fmt.Println("vers length 2")
		version.Tags = vers[0]
	}

	fmt.Printf("\nVers: %s\n", vers)
	fmt.Printf("version Long: %v\n", version.Long)
	fmt.Printf("\nGit: %s\n", version.Git)
	fmt.Printf("\nTags: %s\n", version.Tags)
	fmt.Printf("\nDirty: %v\n", version.Dirty)
}
