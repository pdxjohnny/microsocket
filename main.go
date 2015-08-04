package main

import (
	"log"

	"github.com/pdxjohnny/microsocket/server"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Println(err)
	}
}
