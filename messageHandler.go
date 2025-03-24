package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
)

func insertLoadInInput(userInput string, ctx context.Context, firestoreClient *firestore.Client) string {
	inputSplit := strings.Split(userInput, "]")
	if len(inputSplit) != 2 {
		return "Error"
	}
	databaseIndex, _ := strconv.Atoi(inputSplit[0][1:])
	loadValue := load(ctx, firestoreClient, databaseIndex)
	loadValueStr := strconv.FormatFloat(loadValue, 'f', -1, 64)
	return loadValueStr + inputSplit[1]
}

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
				lastValue := loadLatest(ctx, firestoreClient)
				lastValueStr := strconv.FormatFloat(lastValue, 'f', -1, 64)
				userInput = lastValueStr + userInput
				result := getResult(userInput)
				saveCalc(userInput, result, writer, firestoreClient, ctx)
				returnMessage = strconv.FormatFloat(result, 'f', -1, 64)
			}
		case firstChar == '[':
			{
				loadedUserInput := insertLoadInInput(userInput, ctx, firestoreClient)
				result := getResult(loadedUserInput)
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

func handleDelRequest(writer http.ResponseWriter, ctx context.Context, client *firestore.Client) {
	deleteHistory(writer, ctx, client)
}

// Handle incoming messages
func messageHandler(firestoreClient *firestore.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			handleGetRequest(writer, request, firestoreClient)
		case http.MethodDelete:
			handleDelRequest(writer, request.Context(), firestoreClient)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}

	}
}
