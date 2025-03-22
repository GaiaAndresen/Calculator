package main

import "testing"

func TestEmptyHist(t *testing.T) {
	exp := 0
	res := len(records)
	if exp != res {
		t.Errorf("Expected %d but got %d", exp, res)
	}
}

// These tests are saved to show the testing of a previous version

/*
func TestSave(t *testing.T) {
	inputStr := "1+1"
	val := 2.0
	lenExp := 1
	saveCalc(inputStr, val)

	if lenExp != len(records) {
		t.Errorf("Expected %d elements in history but got %d", lenExp, len(records))
	}

	if inputStr != records[0].Userinput {
		t.Errorf("Expected the input %s to be saved but got %s", inputStr, records[0].Userinput)
	}

	if val != records[0].Result {
		t.Errorf("Expected the result %f to be saved but got %f", val, records[0].Result)
	}
}

func TestSaveMultiple(t *testing.T) {
	inputStr := "1+1"
	val := 2.0
	lenExp := 2
	saveCalc(inputStr, val)
	saveCalc(inputStr, val)

	if lenExp != len(records) {
		t.Errorf("Expected %d elements in history but got %d", lenExp, len(records))
	}
}

func TestClearMem(t *testing.T) {
	inputStr := "1+1"
	val := 2.0
	saveCalc(inputStr, val)
	clearMemory()
	lenExp := 0

	if lenExp != len(records) {
		t.Errorf("Expected %d elements in history after clear but got %d", lenExp, len(records))
	}
}

func TestLoad(t *testing.T) {
	inputStr := "1+1"
	val := 2.0
	saveCalc(inputStr, val)
	loadRec := load(0)

	if inputStr != loadRec.Userinput {
		t.Errorf("Expected the input %s to be loaded but got %s", inputStr, loadRec.Userinput)
	}

	if val != loadRec.Result {
		t.Errorf("Expected the result %f to be loaded but got %f", val, loadRec.Result)
	}
}

func TestLoadFromStrEmpty(t *testing.T) {
	inputStr := "1+1"
	val := 2.0
	saveCalc(inputStr, val)
	loadRes := loadResFromString("")

	if val != loadRes {
		t.Errorf("Expected the result %f to be loaded but got %f", val, loadRes)
	}
}

func TestLoadFromStrEmptyMultiple(t *testing.T) {
	inputStr1 := "1+1"
	val1 := 2.0
	inputStr2 := "1*1"
	val2 := 1.0
	saveCalc(inputStr1, val1)
	saveCalc(inputStr2, val2)
	loadRes := loadResFromString("")

	if val2 != loadRes {
		t.Errorf("Expected the result %f to be loaded but got %f", val2, loadRes)
	}
}

func TestLoadFromStrIndex(t *testing.T) {
	inputStr1 := "1+1"
	val1 := 2.0
	inputStr2 := "1*1"
	val2 := 1.0
	inputStr3 := "1-1"
	val3 := 0.0
	saveCalc(inputStr1, val1)
	saveCalc(inputStr2, val2)
	saveCalc(inputStr3, val3)
	loadRes := loadResFromString("1")

	if val2 != loadRes {
		t.Errorf("Expected the result %f to be loaded but got %f", val2, loadRes)
	}
}

func TestLoadFromStrIndexNoise(t *testing.T) {
	inputStr1 := "1+1"
	val1 := 2.0
	inputStr2 := "1*1"
	val2 := 1.0
	inputStr3 := "1-1"
	val3 := 0.0
	saveCalc(inputStr1, val1)
	saveCalc(inputStr2, val2)
	saveCalc(inputStr3, val3)
	loadRes := loadResFromString("  (1 )")

	if val2 != loadRes {
		t.Errorf("Expected the result %f to be loaded but got %f", val2, loadRes)
	}
}

func TestLoadFromStrErr(t *testing.T) {
	inputStr1 := "1+1"
	val1 := 2.0
	inputStr2 := "1*1"
	val2 := 1.0
	saveCalc(inputStr1, val1)
	saveCalc(inputStr2, val2)
	loadRes := loadResFromString("NotValid")
	errRes := 0.0

	if errRes != loadRes {
		t.Errorf("Expected the error result %f to be loaded but got %f", errRes, loadRes)
	}
}

func TestSavedInParser(t *testing.T) {
	inputStr1 := "1+1"
	inputStr2 := "1*1"
	val1 := 2.0
	getResult(inputStr1)
	getResult(inputStr2)
	loadRes := loadResFromString("0")
	if val1 != loadRes {
		t.Errorf("Expected the result %f to be loaded but got %f", val1, loadRes)
	}
}
*/
