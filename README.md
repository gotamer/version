[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go v1.9](https://img.shields.io/badge/Go-v1.9-green.svg)](http://golang.org)
[![Go Report Card](https://goreportcard.com/badge/bitbucket.org/gotamer/version)](https://goreportcard.com/report/bitbucket.org/gotamer/version)
[![GoDoc](https://godoc.org/bitbucket.org/gotamer/version?status.svg)](https://godoc.org/bitbucket.org/gotamer/version)

# Package version gets version information from git

This works well for simple applications  
It just runs ´git describe --always --long --dirty´ and formats that info into the Version struct
