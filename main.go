package main

import (
	"flag"
	"log"
	"os"
	"github.com/zianKazi/ner-4-reddit/reddit"
)

func main() {
	// TODO: Use config file
	usernamePtr := flag.String("username", "", "Username")
	passwordPtr := flag.String("password", "", "Password")
	clientIdPtr := flag.String("client-id", "", "Client Id")
	clientSecretPtr := flag.String("client-secret", "", "Client Secret")
	flag.Parse()

	log.Println("Application has started with the following flags...")
	log.Println("username: ", *usernamePtr)
	log.Println("password: ", *passwordPtr)
	log.Println("client-id: ", *clientIdPtr)
	log.Println("client-secret: ", *clientSecretPtr)

	if *usernamePtr == "" || *passwordPtr == "" || *clientIdPtr == "" || *clientSecretPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Println("Additional arg size: ", len(flag.Args()))
	for i, flag := range flag.Args() {
		log.Printf("arg index: %d, arg value: %s", i, flag)
	}

	// Initialize Reddit client
	client, clientInitError := reddit.InitClient(&reddit.Credentials{
		Username:     *usernamePtr,
		Password:     *passwordPtr,
		ClientId:     *clientIdPtr,
		ClientSecret: *clientSecretPtr,
	})
	if clientInitError != nil {
		log.Fatal("Failed to initialize Reddit client %w", clientInitError)
	}

	log.Printf("Is client initialized? %t", client.Initialized)

}
