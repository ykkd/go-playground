// Ref: https://docs.google.com/presentation/d/1DtWB-8FcnNb9asxSpIaOLYbAEc9OjBAwMLNxKnPA8pc/edit#slide=id.g4cbe4d134e_0_0

package main

import (
	"fmt"
)

func main() {
	fmt.Println("chapter 3")
	try1()
}

func try1() {
	var sum int
	sum = 5 + 6 + 3 + 20
	avg := sum / 3

	if float32(avg) > 4.5 {
		fmt.Println("good")
	}
}
