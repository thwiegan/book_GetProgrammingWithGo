package main

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"


func TestCalculateWeightAndAgeSingleStatement(t *testing.T) {

	t.Log("Given the need to test the single statement weight and age calculator.")

	if buf, err := CalculateWeightAndAgeSingleStatement(); err != nil {
		t.Fatal("\tShould throw no error.", ballotX, err)
	} else {
		t.Log("\tShould throw no error.", checkMark)

		if (buf.String() != "My weight on the surface of Mars is 24.2112 kg, and I would be 21 years old.") {
			t.Fatal("Should print the correct weight and age.", ballotX, buf.String())
		} else {
			t.Log("Should print the correct weight and age.", checkMark)
		}
	}
}

func TestCalculateWeightAndAgeMutliStatement(t *testing.T) {

	t.Log("Given the need to test the multi statement weight and age calculator.")

	if buf, err := CalculateWeightAndAgeMultiStatement(); err != nil {
		t.Fatal("\tShould throw no error.", ballotX, err)
	} else {
		t.Log("\tShould throw no error.", checkMark)

		if (buf.String() != "My weight on the surface of Mars is 24.2112 kg, and I would be 21 years old.") {
			t.Fatal("Should print the correct weight and age.", ballotX, buf.String())
		} else {
			t.Log("Should print the correct weight and age.", checkMark)
		}
	}
}

func TestCalculateTravelTime(t *testing.T) {
	t.Log("Given the need to test the travel time calculator")

	if buf, err := CalculateTravelTime(); err != nil {
		t.Fatal("\tShould throw no error", ballotX, err)
	} else {
		t.Log("\tShould throw no error.", checkMark)

		if (buf.String() != "186 seconds\n1337 seconds\n39 days\n") {
			t.Fatal("\tShould calculate the correct time.", ballotX, buf.String())
		} else {
			t.Log("\tShould calculate the correct time.", checkMark)
		}
	}
}

func BenchmarkCalculateWeightAndAgeSingleStatement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateWeightAndAgeSingleStatement()
	}
}

func BenchmarkCalculateWeightAndAgeMultiStatement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateWeightAndAgeMultiStatement()
	}
}

func BenchmarkCalculateTravelTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTravelTime()
	}
}