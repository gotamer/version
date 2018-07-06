// Package version gets version information from git
//
// This works well for simple applications
// It just runs ´git describe --always --long --dirty´ and formats
// that info into the Version struct
// Then it creates a file called versioninfo.go
//
//
// cat versioninfo.go
/*
   package main
   const VerLong = 0.2-6-ge29c9a3-dirty

   const VerTag = 0.2
   const VerGit = ge29c9a3
   const VerModTime = 2018-07-06 13:37:03.28865131 +0300 +03
*/
package version
