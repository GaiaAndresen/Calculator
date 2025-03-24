package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Message struct to represent the data
type Message struct {
	Sender    string `json:"sender"` // Could be a user ID or name
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"` // Unix timestamp
}

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

	var loadedValue float64
	loadedValueExists := false
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if len(input) > 0 {
			firstChar := input[0]
			switch {
			case input == "":
			case input == "quit" || input == "q":
				{
					fmt.Println("Bye")
					return
				}
			case input == "clear":
				{
					clearMemory()
					loadedValueExists = false
					fmt.Println("History cleared")
				}
			case len(input) >= 3 && input[:3] == "del":
				{
					loadedValueExists = false
				}
			/*case len(input) >= 4 && input[:4] == "hist":
			{
				printHistory()
			}*/
			case (firstChar >= '0' && firstChar <= '9') || firstChar == '(' || firstChar == '-':
				{
					result := getResult(input)
					fmt.Println(result)
				}
			case firstChar == '+' || firstChar == '-' || firstChar == '*' || firstChar == '/' || firstChar == '^':
				{
					if loadedValueExists {
						input = strconv.FormatFloat(loadedValue, 'f', -1, 64) + input
						result := getResult(input)
						fmt.Println(result)
					} else {
					}
				}
			default:
				{
					fmt.Println("Please enter a valid command or expression")
				}
			}
		}
	}
}
