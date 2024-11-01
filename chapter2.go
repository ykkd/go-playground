// Ref: https://docs.google.com/presentation/d/1CIMDenDLZ7NPNgzmfbCNH_W3dYjaTEBdUYfUuXXuMHk/edit#slide=id.g4e29971f9a_0_0

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// MARK: -  変数

var n int = 100

func main() {
	omikuji2()
}

func variables() {
	// 変数宣言と代入を分ける（関数外＝パッケージレベルでは不可）
	var n2 int
	n2 = 100
	fmt.Println(n2) // n2を出力することで使用する

	// varの省略（関数ないのみ）
	l := 100
	fmt.Println(l)

	// 宣言をまとめる
	var (
		a = 10
		b = "test"
	)

	fmt.Println(a)
	fmt.Println(b)

}

// 型推論
var n3 = 100

// MARK: - 定数
func constants() {
	const a string = "constant"
	const b = 100

	// ConstSpecはGoの定数宣言のための仕様（Specification）で、定数を定義するconstキーワードを使って複数の定数を一度に宣言できる構文
	const (
		c = 200
		d = 300
	)

	const e = c + d

	fmt.Println(e)

	// 定数は型を持たないので無限の精度になる（実装上の制約はある）
	const n = 10000000 / 1000000
	// 型なし定数は案目的に型変換されるので、キャストする手間を省くことができる
	const h = n * 10
	fmt.Println(h)

	// 右辺の省略: 名前付き定数定義の右辺が省略できる
	const (
		x = 1 + 2
		y
		z
	)
	fmt.Println(x, y, z)

	// iota: 定数の列挙, 0から始まるindex的な扱い
	const (
		A = iota // 0
		B        // 1
		C        // 2
	)

	const (
		E = 1 << iota // 1 (2^0)
		F             // 2 (2^1)
		G             // 4 (2^2)
	)
}

// MARK: - 制御構文
func controlSyntax() {
	ifStatement()
	switchStatement()
	loopStatement()
}

func ifStatement() {
	x := 1
	if x == 1 {
		fmt.Println("x is 1")
	} else if x != 2 {
		fmt.Println("x is not 1 or 2")
	} else {
		fmt.Println("x is 2")
	}
}

func switchStatement() {
	a := 2

	switch a {
	case 1, 2:
		fmt.Println("a is 1 or 2")
	default:
		fmt.Println("default")
	}

	// trueを省略可能
	// 何もしない場合は案目的にbreakになるので、全条件をカバーする必要がない
	// caseを跨ぐ際はfallthroughを用いる
	switch {
	case a == 1:
		fmt.Println("a is 1")
	}
}

func loopStatement() {
	// 繰り返しはforしかない

	// 初期値;継続条件;更新文
	for i := 0; i <= 100; i = i + 1 {
		fmt.Println(i)
	}

	// for {
	// 無限ループ
	// }

	// rangeを用いた繰り返し
	for j := range []int{1, 2, 3} {
		fmt.Println(j)
	}
}

func oddOrEven() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("%d-偶数\n", i)
		} else {
			fmt.Println("%d-奇数\n", i)
		}
	}

	// switch版
	for i := 1; i <= 100; i++ {
		switch {
		case i%2 == 0:
			fmt.Printf("%d-偶数\n", i)
		default:
			fmt.Printf("%d-奇数\n", i)
		}
	}
}

// MARK: - 演習

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
