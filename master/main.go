package main

import (
	"niffler/master/fortune"

	"github.com/mysll/toolkit"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	go fortune.Serv()
	toolkit.WaitForQuit()
}
