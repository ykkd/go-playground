package main

import (
	"fmt"
	"math/rand"
	"time"
)

func omikuji() {
	array := []int{1, 2, 3, 4, 5, 6}
	randomIndex := randomValue(array)
	selectedNumber := array[randomIndex]

	switch selectedNumber {
	case 1:
		fmt.Println("凶")
	case 2, 3:
		fmt.Println("吉")
	case 4, 5:
		fmt.Println("中吉")
	case 6:
		fmt.Println("大吉")
	}
}

func randomValue(array []int) int {
	t := time.Now().UnixNano()
	src := rand.NewSource(t)
	r := rand.New(src)
	randomNumber := r.Intn(len(array))
	return randomNumber
}

func omikuji2() {
	r := randomGenerator()
	randomNumber := r.Intn(6) // 0-5

	switch randomNumber + 1 {
	case 1:
		fmt.Println("凶")
	case 2, 3:
		fmt.Println("吉")
	case 4, 5:
		fmt.Println("中吉")
	case 6:
		fmt.Println("大吉")
	}
}

func randomGenerator() *rand.Rand {
	t := time.Now().UnixNano()
	src := rand.NewSource(t)
	r := rand.New(src)
	return r
}
