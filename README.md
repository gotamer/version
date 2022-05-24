[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go v1.9](https://img.shields.io/badge/Go-v1.9-green.svg)](http://golang.org)
[![Go Report Card](https://goreportcard.com/badge/bitbucket.org/gotamer/version)](https://goreportcard.com/report/github.com/gotamer/version)
[![GoDoc](https://godoc.org/github.com/gotamer/version?status.svg)](https://godoc.org/github.com/gotamer/version)


# Package version gets version information from git

```
   ______    ______                         _    __               _           
  / ____/___/_  __/___ _____ ___  ___  ____| |  / /__  __________(_)___  ____ 
 / / __/ __ \/ / / __ `/ __ `__ \/ _ \/ ___/ | / / _ \/ ___/ ___/ / __ \/ __ \
/ /_/ / /_/ / / / /_/ / / / / / /  __/ /   | |/ /  __/ /  (__  ) / /_/ / / / /
\____/\____/_/  \__,_/_/ /_/ /_/\___/_/    |___/\___/_/  /____/_/\____/_/ /_/ 

```

It runs `git describe --always --long --tags --dirty` and formats that info in to a file called version_info.go (see output below)

```go
package main

import "github.com/gotamer/version"

func init() {
	if err := version.Run(); err != nil {
		// "Handle error"
	}
}
```

### cat version_info.go

```go
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

```
