package main

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestRandomNumber(t *testing.T) {

	table := []struct {
		from 	int
		to 	int
	} {
		{
			0,
			10,
		},
		{
			1000,
			10000,
		},
		{
			56000000,
			401000000,
		},
	}

	t.Log("Given the need to test the random number generator.")

	for _, data := range table {
		num, err := RandomNumber(data.from, data.to)

		if (err != nil) {
			t.Log("\tShould throw no error.", ballotX, err)
		} else {
			t.Log("\tShould throw no error.", checkMark)
		}

		if (num < data.from || num > data.to) {
			t.Logf("\t%v should be inbetween %v and %v. %v", num, data.from, data.to, ballotX)
		} else {
			t.Logf("\t%v should be inbetween %v and %v. %v", num, data.from, data.to, checkMark)
		}
	}

}

func BenchmarkRandomNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomNumber(0, 1000000)
	}
}