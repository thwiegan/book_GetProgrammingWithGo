package main

import (
	"fmt"
	"bytes"
)

func main() {
	buf, _ := CalculateWeightAndAgeSingleStatement()
	fmt.Println(buf.String())

	buf, _ = CalculateTravelTime()
	fmt.Println(buf.String())

	speed, _ := HowFastToMalacandra()
	fmt.Println(speed, "km/h")
}

func CalculateWeightAndAgeSingleStatement() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	fmt.Fprint(buf, "My weight on the surface of Mars is ", 64 * 0.3783, " kg, and I would be ", 40 * 365 / 687, " years old.")

	return buf, nil
}

func CalculateWeightAndAgeMultiStatement() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	fmt.Fprint(buf, "My weight on the surface of Mars is ")
	fmt.Fprint(buf, 64 * 0.3783)
	fmt.Fprint(buf, " kg, and I would be ")
	fmt.Fprint(buf, 40 * 365 / 687)
	fmt.Fprint(buf, " years old.")

	return buf, nil
}

func CalculateTravelTime() (*bytes.Buffer, error) {
	const (
		hoursPerDay = 24
		lightSpeed = 299792 // km/s
		spaceXSpeed = 100800 // km/h
	)

	distance := 56000000 // km

	buf := &bytes.Buffer{}

	fmt.Fprintln(buf, distance / lightSpeed, "seconds")

	distance = 401000000
	fmt.Fprintln(buf, distance / lightSpeed, "seconds")

	distance = 96300000
	fmt.Fprintln(buf, distance / spaceXSpeed / hoursPerDay, "days")

	return buf, nil
}

func HowFastToMalacandra() (int, error) {
	const(
		travelDays = 28 // days
		distance = 56000000
		hoursPerDay = 24
	)

	speed := distance / (travelDays * hoursPerDay)

	return speed, nil
}