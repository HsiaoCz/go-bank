package main

import (
	"log"
)

func main() {
	store, err := NewMysqlStore()
	if err != nil {
		log.Fatal(err)
	}
	server := NewApiServer(":9090", store)
	server.Run()
}
