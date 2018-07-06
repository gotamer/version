[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go v1.9](https://img.shields.io/badge/Go-v1.9-green.svg)](http://golang.org)
[![Go Report Card](https://goreportcard.com/badge/bitbucket.org/gotamer/version)](https://goreportcard.com/report/bitbucket.org/gotamer/version)
[![GoDoc](https://godoc.org/bitbucket.org/gotamer/version?status.svg)](https://godoc.org/bitbucket.org/gotamer/version)


# Package version gets version information from git

It just runs ´git describe --always --long --dirty´ and formats that info into the Version struct.
Then it creates a file called versioninfo.go

### versioninfo.go
´´´go
package main
const VerLong = 0.2-6-ge29c9a3-dirty

const VerTag = 0.2
const VerGit = ge29c9a3
const VerModTime = 2018-07-06 13:37:03.28865131 +0300 +03
´´´