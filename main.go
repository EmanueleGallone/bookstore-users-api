package main

import (
	"bookstore-users-api/app"
	"log"
)

func main() {
	//if go code crashes, this returns the filename and the line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app.StartApplication()
}
