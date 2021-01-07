package main

import (
	"flag"
	"github.com/zianKazi/ner-4-reddit/reddit"
	"log"
	"os"
)

func main() {
	// TODO: Use config file
	username := flag.String("username", "", "Username")
	password := flag.String("password", "", "Password")
	clientId := flag.String("client-id", "", "Client Id")
	clientSecret := flag.String("client-secret", "", "Client Secret")
	subreddit := flag.String("subreddit", "","Sub-reddit")
	flag.Parse()

	log.Println("Application has started with the following flags...")
	log.Println("username: ", *username)
	log.Println("password: ", *password)
	log.Println("client-id: ", *clientId)
	log.Println("client-secret: ", *clientSecret)
	log.Println("subreddit: ", *subreddit)

	if *username == "" || *password == "" || *clientId == "" || *clientSecret == "" || *subreddit == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Println("Additional arg size: ", len(flag.Args()))
	for i, flag := range flag.Args() {
		log.Printf("arg index: %d, arg value: %s", i, flag)
	}

	// Initialize Reddit client
	client, clientInitError := reddit.InitClient(&reddit.Credentials{
		Username:     *username,
		Password:     *password,
		ClientId:     *clientId,
		ClientSecret: *clientSecret,
	})
	if clientInitError != nil {
		log.Fatal("Failed to initialize Reddit client %w", clientInitError)
	}

	log.Printf("Is client initialized? %t", client.Initialized)
	responseChannel := make(chan *reddit.ResponseData)
	errorChannel := make(chan error)
	go client.Read(*subreddit, responseChannel, errorChannel)
	for {
		select {
			case data := <- responseChannel:
				log.Println("Response received")
				log.Printf("%d", len(data.Data.Children))
				return
			case error:= <-	errorChannel:
				log.Println("Error occurred")
				log.Fatal(error)
		}
	}

}
