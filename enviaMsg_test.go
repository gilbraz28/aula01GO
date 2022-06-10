package main

import (
	"fmt"
	"testing"
)

// - struct test, fields: data []int, answer int
// - tests := []test{[]int{}, int}
// - range tests

type test struct {
	data   []int
	answer int
}

var tests = []test{
	{data: []int{1, 2}, answer: 3},
	{[]int{10, 11}, 21},
	{[]int{-5, 15}, 10},
}

func ExampleExecSoma() {
	fmt.Println(execSoma(5, 5))
	fmt.Println(execSoma(25, 25))
	// Output:
	// 10
	// 50
}

func TestSomaEmTabela(t *testing.T) {

	for _, v := range tests {
		z := execSoma(v.data[0], v.data[1])
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got:", z)
		}
	}
}

func TestExecSoma(t *testing.T) {
	teste := execSoma(5, 5)
	resultado := 10
	if teste != resultado {
		t.Error("Expected: ", resultado, "Got: ", teste)
	}
}

func TestExecMultiplica(t *testing.T) {
	teste := execMultiplicacao(5, 5)
	resultado := 25
	if teste != resultado {
		t.Error("Expected: ", resultado, "Got: ", teste)
	}
}
