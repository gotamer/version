package version

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// Version holds version info from git describe
var Version struct {
	Long    string
	Commits int
	Tags    string
	Git     string
	Dirty   bool
}

func init() {
	version()
}

// Version gets version info from git describe
func version() {
	var err error
	cmd := exec.Command("git", "describe", "--always", "--long", "--tags", "--dirty")
	cmd.Stdin = strings.NewReader("Version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Printf("Version : Error: %s", err.Error())
	} else {
		Version.Long = out.String()
		vers := strings.Split(Version.Long, "-")
		l := len(vers)
		if strings.Contains(vers[l-1], "dirty") {
			Version.Dirty = true
			vers = vers[0 : l-1]
		}
		l = len(vers)
		Version.Git = vers[l-1]
		vers = vers[0 : l-1]
		l = len(vers)

		if len(vers) == 2 {
			Version.Tags = vers[0]
			if Version.Commits, err = strconv.Atoi(vers[1]); err != nil {
				log.Printf("Version : Error: %s", err.Error())
			}
		}
	}
}
