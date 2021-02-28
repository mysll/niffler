package main

import (
	"niffler/master/fortune"

	"github.com/mysll/toolkit"
)

func main() {
	go fortune.Serv()
	toolkit.WaitForQuit()
}
