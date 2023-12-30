package main

import "github.com/q-sw/go-bank-analysis/nordigen-api-client/cmd"

func main() {
	server := cmd.NewServer(":8080")
	server.Run()
}
