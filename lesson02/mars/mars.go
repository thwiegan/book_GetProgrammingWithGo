package main

import (
	"fmt"
	"bytes"
)

func main() {
	buf, _ := CalculateSingleStatement()
	fmt.Print(buf.String())
}

func CalculateSingleStatement() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	fmt.Fprint(buf, "My weight on the surface of Mars is ", 64 * 0.3783, " kg, and I would be ", 40 * 365 / 687, " years old.")

	return buf, nil
}

func CalculateMultiStatement() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	fmt.Fprint(buf, "My weight on the surface of Mars is ")
	fmt.Fprint(buf, 64 * 0.3783)
	fmt.Fprint(buf, " kg, and I would be ")
	fmt.Fprint(buf, 40 * 365 / 687)
	fmt.Fprint(buf, " years old.")

	return buf, nil
}