package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
)

type Record struct {
	Userinput string
	Result    float64
	Time      int64
}

var records []Record = make([]Record, 0)

func saveCalc(userInput string, res float64, writer http.ResponseWriter, firestoreClient *firestore.Client, ctx context.Context) {
	records = append(records, Record{Userinput: userInput, Result: res})
	record := Record{Userinput: userInput, Result: res, Time: time.Now().Unix()}

	_, _, err := firestoreClient.Collection("calculations").Add(ctx, record)
	if err != nil {
		log.Printf("Failed to add message: %v", err)
		http.Error(writer, "Failed to store message", http.StatusInternalServerError)
		return
	}
}

func getHistoryString(firestoreClient *firestore.Client, ctx context.Context) string {
	// Query Firestore to get all records
	iter := firestoreClient.Collection("calculations").OrderBy("Time", firestore.Asc).Documents(ctx)
	defer iter.Stop()

	history := ""

	for i := 0; ; i++ {
		doc, err := iter.Next()
		if err == io.EOF || doc == nil {
			break
		}
		if err != nil {
			fmt.Println(err)
			return "Could not load history, error with iterator"
		}

		var record Record
		if err := doc.DataTo(&record); err != nil {
			return "Could not load history, error with load of record"
		}
		history = history + "[" + strconv.Itoa(i) + "]  " + record.Userinput + " = " + strconv.FormatFloat(record.Result, 'f', -1, 64) + "\n"
	}
	return history
}

func clearMemory() {
	records = make([]Record, 0)
}

func loadLatest() {
	//Limit(1)
}

func load(index int) Record {
	return records[index]
}

func loadResFromString(numStr string) float64 {
	numStr = strings.ReplaceAll(numStr, " ", "")
	numStr = strings.ReplaceAll(numStr, "(", "")
	numStr = strings.ReplaceAll(numStr, ")", "")
	recordsAmount := len(records)
	if numStr == "" {
		if recordsAmount == 0 {
			fmt.Println("No values in memory, insert 0")
			return 0
		}
		return load(recordsAmount - 1).Result
	}
	number, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error, expected an index, but got:", numStr)
		return 0
	}
	if number >= recordsAmount {
		fmt.Printf("There are only %d elements in memory, (%d) is invalid", recordsAmount, number)
		return 0
	}
	return load(int(number)).Result
}
