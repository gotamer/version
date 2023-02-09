[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go v1.9](https://img.shields.io/badge/Go-v1.9-green.svg)](http://golang.org)
[![Go Report Card](https://goreportcard.com/badge/bitbucket.org/gotamer/version)](https://goreportcard.com/report/github.com/gotamer/version)
[![GoDoc](https://godoc.org/github.com/gotamer/version?status.svg)](https://godoc.org/github.com/gotamer/version)


# Package version gets version information from git to go

```
 ..|'''.|   ||    .         ..|'''.|
.|'     '  ...  .||.   /\  .|'     '    ...
||    ....  ||   ||   (  ) ||    .... .|  '|.
'|.    ||   ||   ||     // '|.    ||  ||   ||
 ''|...'|  .||.  '|.'  //   ''|...'|   '|..|'
                      /(
                      {___

'||'  '|'                         ||
 '|.  .'    ....  ... ..   ....  ...    ...   .. ...
  ||  |   .|...||  ||' '' ||. '   ||  .|  '|.  ||  ||
   |||    ||       ||     . '|..  ||  ||   ||  ||  ||
    |      '|...' .||.    |'..|' .||.  '|..|' .||. ||.

FREE THOUGHT · FREE SOFTWARE · FREE WORLD
```

Version runs `git describe --always --long --tags --dirty` and formats that info
by default to stdout for you to view or pipe, or optionally in to a file (version.go)
(see sample output below)

If you rather not include another package in your app then see [verup](https://github.com/gotamer/version/tree/master/verup).
[verup](https://github.com/gotamer/version/tree/master/verup) is a cmd you can execute at your main/cmd apps folder and it will do the same thing.


### $ cat version.go
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

### Alternative:
https://belief-driven-design.com/build-time-variables-in-go-51439b26ef9
