package main

import (
	"log"

	"github.com/gotamer/version"
)

func main() {
	var err = version.Run()
	if err != nil {
		log.Println("ERR Version : ", err.Error())
	}
}
