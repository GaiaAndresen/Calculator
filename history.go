package main

import (
	"fmt"
	"time"
)

type Record struct {
	userinput string
	result    float64
}

var records []Record = make([]Record, 0)

func getTime() time.Time {
	return time.Now()
}

func saveCalc(userinput string, res float64) {
	records = append(records, Record{userinput: userinput, result: res})
}

func printHistory() {
	for i, rec := range records {
		fmt.Printf("(%d): %s = %.2f\n", i, rec.userinput, rec.result)
	}
}

func clearMemory() {
	records = make([]Record, 0)
}

func load(num string) float64 {
	return 1
}

//Use last result
//See all
//Load mem number
//Clear mem
