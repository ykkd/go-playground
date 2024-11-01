// Ref: https://docs.google.com/presentation/d/1DtWB-8FcnNb9asxSpIaOLYbAEc9OjBAwMLNxKnPA8pc/edit#slide=id.g4cbe4d134e_0_0

package main

import (
	"fmt"
)

func main() {
	fmt.Println("chapter 3")
	try1()
	composite1()
}

func try1() {
	var sum int
	sum = 5 + 6 + 3 + 20
	avg := sum / 3

	if float32(avg) > 4.5 {
		fmt.Println("good")
	}
}

// MARK: - コンポジット型
func composite1() {
	// 1. 要素がint型であるスライス。初期値はnil
	var ns []int
	fmt.Println(ns) // []

	// 2. map型(key: string, value: int) 初期値はnil
	var m map[string]int
	fmt.Println(m) // map[]

	// 3. 構造体。初期値は全てゼロ値
	var p struct {
		name string
		age  int
	}

	fmt.Println(p) // { 0}

	// 4. 配列 要素の型は同じ。要素数を指定する必要があり、かつ、要素数が違えば別の型。要素数を指定しないとスライス型になる
	var array [5]int
	fmt.Println(array) // 0 0 0 0 0
}
