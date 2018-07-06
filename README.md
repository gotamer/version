[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go v1.9](https://img.shields.io/badge/Go-v1.9-green.svg)](http://golang.org)
[![Go Report Card](https://goreportcard.com/badge/bitbucket.org/gotamer/version)](https://goreportcard.com/report/bitbucket.org/gotamer/version)
[![GoDoc](https://godoc.org/bitbucket.org/gotamer/version?status.svg)](https://godoc.org/bitbucket.org/gotamer/version)


# Package version gets version information from git

It just runs ´git describe --always --long --dirty´ and formats that info into the Version struct.
Then it creates a file called versioninfo.go

### versioninfo.go
```go
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
```