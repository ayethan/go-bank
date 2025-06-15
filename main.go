package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fmt.Printf("%+v\n", store)

	// server := NewAPIServer(":3000", store)
	// server.run()
}
