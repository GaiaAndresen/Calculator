package main

import (
	"math"
	"testing"
)

func TestInt(t *testing.T) {
	input := "1"
	var exp float64 = 1
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestFloat(t *testing.T) {
	input := "1.2"
	var exp float64 = 1.2
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestAdd(t *testing.T) {
	input := "1+2"
	var exp float64 = 3
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestSub(t *testing.T) {
	input := "2-1"
	var exp float64 = 1
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestSubNeg(t *testing.T) {
	input := "1-2"
	var exp float64 = -1
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestMult(t *testing.T) {
	input := "2*3"
	var exp float64 = 6
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestDiv(t *testing.T) {
	input := "6/3"
	var exp float64 = 2
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestDivFloat(t *testing.T) {
	input := "1/2"
	var exp float64 = 0.5
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestExp(t *testing.T) {
	input := "2^3"
	var exp float64 = 8
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestPar(t *testing.T) {
	input := "(1)"
	var exp float64 = 1
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestParNested(t *testing.T) {
	input := "((1))"
	var exp float64 = 1
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestParParallel(t *testing.T) {
	input := "(1)+(1)"
	var exp float64 = 2
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestMinLeft(t *testing.T) {
	input := "4-2-1"
	var exp float64 = 1
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestDivLeft(t *testing.T) {
	input := "8/4/2"
	var exp float64 = 1
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestOrderOfOperations(t *testing.T) {
	input := "1+2*3^4*5+6+7*8"
	var exp float64 = 1 + (2 * math.Pow(3, 4) * 5) + 6 + (7 * 8)
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}

func TestBig(t *testing.T) {
	input := "(((1+2*3)+5^2)+1.4^1.2*(2+1)^1)/2-(3+5)"
	var exp float64 = ((7+25)+(math.Pow(1.4, 1.2)*3))/2 - 8
	res := getResult(input)
	if exp != res {
		t.Errorf(input, "was expected to be ", exp, "but was", res)
	}
}
