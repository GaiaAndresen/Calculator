package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("firebase.json")
	config := &firebase.Config{
		ProjectID: "calculatorfirebase-30c4e",
	}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		fmt.Println("Error creating new app:", err)
	}

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("Error creating Firestore client:", err)
	}
	defer firestoreClient.Close()

	messageHandler := messageHandler(firestoreClient)

	// Register the message handler
	http.Handle("/calculations", messageHandler)

	// Start the server
	fmt.Println("Calculator on")
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
