package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
)

func handleGetRequest(writer http.ResponseWriter, request *http.Request, firestoreClient *firestore.Client) {

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}
	defer request.Body.Close() // Important: Close the request body
	userInput := string(body)
	ctx := request.Context()

	returnMessage := "Empty message received"
	if len(userInput) > 0 {
		firstChar := userInput[0]
		switch {
		case len(userInput) >= 4 && userInput[:4] == "hist":
			{
				fmt.Fprintln(writer, getHistoryString(firestoreClient, ctx))
				return
			}
		case (firstChar >= '0' && firstChar <= '9') || firstChar == '(' || firstChar == '-':
			{
				result := getResult(userInput)
				saveCalc(userInput, result, writer, firestoreClient, ctx)
				returnMessage = strconv.FormatFloat(result, 'f', -1, 64)
			}
		case firstChar == '+' || firstChar == '-' || firstChar == '*' || firstChar == '/' || firstChar == '^':
			{
				lastValue := strconv.FormatFloat(loadResFromString(""), 'f', -1, 64)
				userInput = lastValue + userInput
				result := getResult(userInput)
				saveCalc(userInput, result, writer, firestoreClient, ctx)
				returnMessage = strconv.FormatFloat(result, 'f', -1, 64)
			}
		default:
			{
				returnMessage = "Please enter a valid command or expression"
			}
		}
	}

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, returnMessage)
}

// Handle incoming messages
func messageHandler(firestoreClient *firestore.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			handleGetRequest(writer, request, firestoreClient)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}

	}
}
