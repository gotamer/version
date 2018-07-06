package version

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const gitfile = ".git/COMMIT_EDITMSG"

// Version holds version info from ´git describe´
var Version struct {
	Long    string
	Tags    string
	Git     string
	Dirty   bool
	ModTime time.Time
}

func init() {
	if _, err := os.Stat(gitfile); !os.IsNotExist(err) {
		version()
		modTime()
		write()
	}
}

// Long version from Git
func Long() string {
	return Version.Long
}

// Tag version from Git
func Tag() string {
	return Version.Tags
}

// Git returns the 7 hexadecimal digits version from Git
func Git() string {
	return Version.Tags
}

// ModTime returns the last Git commit time
func ModTime() time.Time {
	return Version.ModTime
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

		if len(vers) == 2 {
			Version.Tags = vers[0]
		}
	}
}

func modTime() {
	fi, err := os.Stat(gitfile)
	if err != nil {
		log.Printf("Version : Error: %s", err.Error())
	}
	Version.ModTime = fi.ModTime()
}

func write() {
	var buf bytes.Buffer
	buf.WriteString("package main\n")
	buf.WriteString(fmt.Sprintf("const VerLong = %s\n", Version.Long))
	buf.WriteString(fmt.Sprintf("const VerTag = %s\n", Version.Tags))
	buf.WriteString(fmt.Sprintf("const VerGit = %s\n", Version.Git))
	buf.WriteString(fmt.Sprintf("const VerModTime = %s\n", Version.ModTime))

	//os.Remove("versioninfo.go")
	f, err := os.Create("versioninfo.go")
	if err != nil {
		log.Printf("Version : file : Error: %s", err.Error())
	}
	defer f.Close()
	buf.WriteTo(f)
}
