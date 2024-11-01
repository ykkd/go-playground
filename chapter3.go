// Ref: https://docs.google.com/presentation/d/1DtWB-8FcnNb9asxSpIaOLYbAEc9OjBAwMLNxKnPA8pc/edit#slide=id.g4cbe4d134e_0_0

package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("chapter 3")
	try2()
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

	// 配列初期化方法
	var array1 [5]int
	fmt.Println(array1) // 0 0 0 0 0

	var array2 = [5]int{10, 20, 30, 40, 50}
	fmt.Println(array2) // 10 20 30 40 50

	array3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(array3) // 10 20 30 40 50

	// index: valueで特定の要素の値を指定。indexは0始まり
	array4 := [...]int{5: 50, 10: 100}
	fmt.Println(array4) // 0, 0, 0, 0, 0, 50, 0, 0, 0, 0, 100

	// 配列の操作
	array5 := [...]int{10, 20, 30, 40, 50}

	fmt.Println(array5[3]) // 40

	fmt.Println(len(array5)) // 5

	fmt.Println(array5[1:2]) // 20 (index1以上、2未満)
}

func slicesExample() {
	ns := []int{10, 20, 30, 40, 50}

	fmt.Println(ns[3])
	fmt.Println(len(ns))
	fmt.Println(cap(ns)) // 容量

	// 要素追加
	ns = append(ns, 60, 70) // appendの挙動　容量が足りない場合は元の配列のおよそ2倍の容量を確保し直す
	fmt.Println(ns, len(ns), cap(ns))

	// capを指定する
	ms := ns[:4:4]
	fmt.Println(ms, len(ms), cap(ms))

	// slice tricks
	ns = append(ns[:3], ns[1:]...)
	fmt.Println(ns)

	ns = slices.Delete(ns, 1, 3)
	fmt.Println(ns)
}

func try2() {
	slice := []int{19, 86, 1, 12}

	var sum int
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	fmt.Println(sum)
}
