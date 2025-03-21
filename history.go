package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Record struct {
	userinput string
	result    float64
}

var records []Record = make([]Record, 0)

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
		return load(recordsAmount - 1).result
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
	return load(int(number)).result
}
