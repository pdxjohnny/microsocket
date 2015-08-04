package main

import (
	"log"

	"github.com/pdxjohnny/mircosocket/server"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Println(err)
	}
}
