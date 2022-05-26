package version

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Run gets version info from git describe
func Run() error {
	fi, err := os.Stat(".git/COMMIT_EDITMSG")
	if os.IsNotExist(err) {
		// Git is not available
		// assuming file is being executed not build
		err = fmt.Errorf("Must be executed in the git root directory")
		return err
	}
	var buf bytes.Buffer
	buf.WriteString("package main\n\n")
	buf.WriteString("//VarModTime is a UTC Unix time stamp\n")
	buf.WriteString(fmt.Sprintf("const VerModTime = %d\n\n", fi.ModTime().UTC().Unix()))

	cmd := exec.Command("git", "describe", "--always", "--long", "--tags", "--dirty")
	cmd.Stdin = strings.NewReader("Version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()

	var long = out.String()
	long = strings.Trim(long, "\n ")

	if err != nil {
		return err
	} else {
		buf.WriteString("//VarLong is the full version from Git command output\n")
		buf.WriteString(fmt.Sprintf("const VerLong = \"%s\"\n\n", long))

		vers := strings.Split(long, "-")
		l := len(vers)
		buf.WriteString("//VarDirty means app was build with a git dir that contained modifications which had not been committed.\n")
		if strings.Contains(vers[l-1], "dirty") {
			buf.WriteString(fmt.Sprintf("const VerDirty = %s\n\n", "true"))
			vers = vers[0 : l-1]
		} else {
			buf.WriteString(fmt.Sprintf("const VerDirty = %s\n\n", "false"))
		}
		l = len(vers)
		buf.WriteString("//VarGit is the 7 hexadecimal digits version from Git.\n")
		buf.WriteString(fmt.Sprintf("const VerGit = \"%s\"\n\n", vers[l-1]))
		vers = vers[0 : l-1]

		if len(vers) == 2 {
			buf.WriteString("//VarTag is the Tag version from Git.\n")
			buf.WriteString(fmt.Sprintf("const VerTag = \"%s\"\n", vers[0]))
		}
	}

	f, err := os.Create("version.go")
	if err == nil {
		buf.WriteTo(f)
		f.Close()
	}
	return err
}