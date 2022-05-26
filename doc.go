// Copyright 2009 The GoTamer. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package version gets version information from git to go

verup will create a file (version.go) with up to date
version information from git.
By default verup will output to stdout to view or pipe.
Please check the flags.


cat version.go

	package main

	//VarModTime is a UTC Unix time stamp
	const VerModTime = 1530896805

	//VarLong is the full version from Git the command output
	const VerLong = "0.2-14-g1051a2c-dirty"

	//VarDirty means app was build with a git dir that contained modifications which had not been committed.
	const VerDirty = true

	//VarGit is the 7 hexadecimal digits version from Git.
	const VerGit = "g1051a2c"

	//VarTag is the Tag version from Git.
	const VerTag = "0.2"

*/
package version
