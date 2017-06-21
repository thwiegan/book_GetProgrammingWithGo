package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	const from = 56000000
	const to = 401000000
	num, _ := RandomNumber(from, to)

	fmt.Printf("Random number between %v and %v: %v", from, to, num)
}


func RandomNumber(from int, to int) (int,error) {
	return rand.Intn(to - from) + from, nil
}