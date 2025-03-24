package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
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

func loadLatest(ctx context.Context, firestoreClient *firestore.Client) float64 {
	query := firestoreClient.Collection("calculations").OrderBy("Time", firestore.Desc).Limit(1)

	iter := query.Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		return 0
	}
	var record Record
	if err := doc.DataTo(&record); err != nil {
		return 0
	}
	return record.Result

}

func load(ctx context.Context, firestoreClient *firestore.Client, index int) float64 {
	query := firestoreClient.Collection("calculations").OrderBy("Time", firestore.Asc).Limit(index + 1)

	iter := query.Documents(ctx)
	defer iter.Stop()

	i := 0
	for {
		doc, err := iter.Next()
		if err == io.EOF || doc == nil {
			return 0
		}
		if err != nil {
			return 0
		}

		if i == index {
			var record Record
			if err := doc.DataTo(&record); err != nil {
				return 0
			}
			return record.Result
		}
		i++
	}
}

func deleteHistory(writer http.ResponseWriter, ctx context.Context, client *firestore.Client) {
	ref := client.Collection("calculations")
	batchSize := 100
	for {
		// Get a batch of documents in the collection
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents and delete them
		writeBatch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				writer.WriteHeader(http.StatusNotFound)
				fmt.Println("Error when iterating")
				fmt.Fprintln(writer, "Could not delete history")
				return
			}

			writeBatch.Delete(doc.Ref) // Use writeBatch.Delete
			numDeleted++
		}

		// If there are no documents to delete, we're done
		if numDeleted == 0 {
			writer.WriteHeader(http.StatusOK)
			fmt.Fprintln(writer, "History deleted")
			return
		}

		// Commit the batch delete
		_, err := writeBatch.Commit(ctx)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			fmt.Println("Error when deleting element")
			fmt.Fprintln(writer, "Could not delete history")
			return
		}
	}
}
