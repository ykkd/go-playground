// Ref: https://docs.google.com/presentation/d/1HW3wG8J_Q2536Iu__7HGr_mhurHajC7IOGjCnn3kZmg/edit#slide=id.g4fa5dcc83a_0_0

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

}

// MARK: - エラー処理

// error  interface
// type error interface {
// 	Error() string
// }

// nilと比較してerrorが発生したかをチェックする
// if err := f(); err != nil {
//     エラー処理
// }

func howToMakeSimpleError() {
	err := errors.New("Error")

	var name string
	err2 := fmt.Errorf("%s is not found", name)

	fmt.Println(err, err2)
}

// MARK: - エラー型の定義

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

// type Stringer interface {
// 	String() string
// }

// func ToStringeer(v interface{}) (Stringer, error) {
// 	if s, ok := v.(Stringer); ok {
// 		return s, nil
// 	}
// 	return nil, MyError("CastError")
// }

// type MyError string

// func (e MyError) Error() string {
// 	return string(e)
// }

// type S string

// func (s S) String() string {
// 	return string(s)
// }

// func execute() {
// 	v := S("Hoge")
// 	if s, err := ToStringer(v); err != nil {
// 		fmt.Fprintln(os.Stderr, "Error:", err)
// 	} else {
// 		fmt.Println(s.String())
// 	}
// }

// MARK: - try

type RuneScanner struct {
	r   io.Reader
	buf [16]byte
}

func NewRuneScanner(r io.Reader) *RuneScanner {
	return &RuneScanner{r: r}
}

func (s *RuneScanner) Scan() (rune, error) {
	n, err := s.r.Read(s.buf[:])
	if err != nil {
		return 0, err
	}

	r, size := utf8.DecodeRune(s.buf[:n])
	if r == utf8.RuneError {
		return 0, errors.New("RuneError")
	}

	s.r = io.MultiReader(bytes.NewReader(s.buf[size:n]), s.r)
	return r, nil
}

func execute() {
	s := NewRuneScanner(strings.NewReader("Hello, 世界"))
	for {
		r, err := s.Scan()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%c\n", r)
	}
}

// error値によって分岐する

func errorSample() {
	errors.Is(err, os.ErrExist) {
		// os.ErrExistだった場合の処理
	}
}

// エラーから情報を取り出す
func getInfoFromError() {
	var pathError *os.PathError

	if errors.As(err, &pathError) {
		fmt.Println("failed at path:", pathError.path)
	} else {
		fmt.Println(err)
	}
}

