package main

import "log"

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		panic(err)
	}
	log.Println("connected to the database : ", store)
	// init database for configuring
	store.Init()

	server := NewAPIServer(":8080", store)
	server.Run()
}
