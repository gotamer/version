// Package version gets version information from git
//
// It runs ´git describe --always --long --tags --dirty´ and formats that info in to a file called versioninfo.go
//
// cat versioninfo.go
/*
	package main

	VarModTime is a UTC Unix time stamp
	const VerModTime = 1530896158

	VarLong is the full version from Git command output
	const VerLong = "0.2-11-g80e8dd8"

	VarDirty means app was build with a git dir that contained modifications which had not been committed.
	const VerDirty = false

	VarGit is the 7 hexadecimal digits version from Git.
	const VerGit = "g80e8dd8"

	VarTag is the Tag version from Git.
	const VerTag = "0.2"
*/
package version
