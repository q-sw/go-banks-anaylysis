package main

import (
	"log"

	"github.com/q-sw/go-bank-analysis/api"
	"github.com/q-sw/go-bank-analysis/storage"
)

func main() {
	str, err := storage.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	if err := str.Init(); err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(":8080")
	server.Run()

}
