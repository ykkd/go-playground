// Ref: https://docs.google.com/presentation/d/1r96GPwlWCkVeeUnFj7dwb8RqMalYLG7vDYWjuAa_Wik/edit#slide=id.g4f15a7d687_0_0
package main

import "fmt"

func main() {
	typeAssertionSample()
}

// Stringer is an interface
type Stringer interface {
	String() string
}

// Hex is a method
type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func interfaceExample() {
	var a Stringer = Hex(100)
	fmt.Println(a)
}

// any type
// any means 'type any interface {}'

func anyInterfaceSample() {
	var v interface{}
	fmt.Println(v)

	var j any
	fmt.Println(j)
}

// 関数にinterfaceを実装する

// Func is a function type which returns string
type Func func() string

func (f Func) String() string { return f() }

func functionInterfaceSample() {
	var s fmt.Stringer = Func(func() string { return "hi" })
	fmt.Println(s)

	// 不可 runtime error
	// var j fmt.Stringer = Func(nil)
	// fmt.Println(j)
}

// type assertion
func typeAssertionSample() {
	var v interface{}

	v = 100

	// {interface}.({type})でインターフェース型を任意の型にキャスト可能。第二戻り値にキャスト可否が返る
	n, ok := v.(int)
	fmt.Println(n, ok)

	s, ok := v.(string)
	fmt.Println(s, ok)

	// switch based on type
	switch i := v.(type) {
	case int:
		fmt.Println(i * 2)
	case string:
		fmt.Println(i + "hoge")
	default:
		fmt.Println("default")
	}
}

// MARK: - 埋め込みとインターフェース

// 構造体の埋め込み: 構造体に匿名フィールドを埋め込む機能

type Hoge struct {
	N int
}

type Fuga struct {
	Hoge // 名前のないフィールドになる
}

func embeddingSample() {
	f := Fuga{Hoge{N: 100}}
	fmt.Println(f.N)      // Hoge型のフィールドにアクセスできる
	fmt.Println(f.Hoge.N) // Hoge型の型名を指定してアクセスすることも可能
}

// インターフェースと埋め込み
// 既存インターフェースの振る舞いを変える

type A interface {
	M()
	N()
}

type B struct{ A } // インターフェースを埋め込む
func (b B) M() { // メソッドを定義する。メソッド名をinterfaceによって提供されている関数名と同じにすることで、インターフェースの振る舞いを変えることが可能
	fmt.Println("hi")
	b.A.M()
}

func hiA(a A) A {
	return B{a}
}
