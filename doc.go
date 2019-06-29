// Package version gets version information from git
//

/*
	package main

	import "bitbucket.org/gotamer/version"

	func init() {
		if err := version.Run(); err != nil {
			// "Handle error"
		}
	}
*/

// It runs ´git describe --always --long --tags --dirty´ and
// formats that info in to a file called version_info.go
//
// cat version_info.go
/*
	package main

	//VarModTime is a UTC Unix time stamp
	const VerModTime = 1530896805

	//VarLong is the full version from Git command output
	const VerLong = "0.2-14-g1051a2c-dirty"

	//VarDirty means app was build with a git dir that contained modifications which had not been committed.
	const VerDirty = true

	//VarGit is the 7 hexadecimal digits version from Git.
	const VerGit = "g1051a2c"

	//VarTag is the Tag version from Git.
	const VerTag = "0.2"

*/
package version
