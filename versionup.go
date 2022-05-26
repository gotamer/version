package version

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Out() error {
	buf, err := getBuf()
	if buf.Len() == 0 {
		// Git is not available
		// assuming file is being executed not build
		return nil
	}

	if err == nil {
		_, err = os.Stdout.Write(buf.Bytes())
	}
	return err
}

func Save(file string) error {
	buf, err := getBuf()
	if buf.Len() == 0 {
		// Git is not available
		// assuming file is being executed not build
		return nil
	}

	if err == nil {
		f, err := os.Create(file)
		if err == nil {
			buf.WriteTo(f)
			f.Close()
		}
	}
	return err
}

func getBuf() (buf bytes.Buffer, err error) {

	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Stdin = strings.NewReader("pwd")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		// Git is not available
		// assuming file is being executed not build
		return buf, nil
	}

	var pwd = strings.Trim(out.String(), "\n ")

	var fi fs.FileInfo
	fi, err = os.Stat(filepath.Join(pwd, ".git/COMMIT_EDITMSG"))
	if os.IsNotExist(err) {
		return
	}

	buf.WriteString("package main\n\n")
	buf.WriteString("//VarModTime is a UTC Unix time stamp\n")
	buf.WriteString(fmt.Sprintf("const VerModTime = %d\n\n", fi.ModTime().UTC().Unix()))

	cmd = exec.Command("git", "describe", "--always", "--long", "--tags", "--dirty")
	cmd.Stdin = strings.NewReader("Version")
	out.Reset()
	cmd.Stdout = &out
	err = cmd.Run()

	var long = out.String()
	long = strings.Trim(long, "\n ")

	if err != nil {
		return
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
	return
}
