package main

import (
	"flag"
	"log"

	"github.com/PrajvalBadiger/go-icecream/api"
	"github.com/PrajvalBadiger/go-icecream/storage"
)

func main() {

	// --listenaddr command line arg, default - :8080
	listen_addr := flag.String("listenaddr", ":8080", "the server address")
	// Parse command line args
	flag.Parse()

	// create storage TODO: can be moved to init()
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	// Initalize a storage
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	// Create a new server
	server := api.New_server(*listen_addr, store)

	// Start the server
	log.Fatal(server.Run())
}
