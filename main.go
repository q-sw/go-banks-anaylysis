package main

import "github.com/q-sw/go-bank-analysis/api"

func main() {
	server := api.NewServer(":8080")
	server.Run()
}
