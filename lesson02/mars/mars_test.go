package main

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"


func TestCalculateSingleStatement(t *testing.T) {

	t.Log("Given the need to test the single statement weight and age calculator.")

	if buf, err := CalculateSingleStatement(); err != nil {
		t.Fatal("\tShould throw no error.", ballotX, err)
	} else {
		t.Log("\tShould throw no error.", checkMark)

		if (buf.String() != "My weight on the surface of Mars is 24.2112 kg, and I would be 21 years old.") {
			t.Fatal("\tShould be 'My weight on the surface of Mars is 24.2112 kg, and I would be 21 years old.'", ballotX, buf.String())
		} else {
			t.Log("\tShould print the correct weight and age.", checkMark)
		}
	}
}

func TestCalculateMultiStatement(t *testing.T) {

	t.Log("Given the need to test the multi statement weight and age calculator.")

	if buf, err := CalculateSingleStatement(); err != nil {
		t.Fatal("\tShould throw no error.", ballotX, err)
	} else {
		t.Log("\tShould throw no error.", checkMark)

		if (buf.String() != "My weight on the surface of Mars is 24.2112 kg, and I would be 21 years old.") {
			t.Fatal("\tShould be 'My weight on the surface of Mars is 24.2112 kg, and I would be 21 years old.'", ballotX, buf.String())
		} else {
			t.Log("\tShould print the correct weight and age.", checkMark)
		}
	}
}

func BenchmarkCalculateSingleStatement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSingleStatement()
	}
}

func BenchmarkCalculateMultiStatement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateMultiStatement()
	}
}