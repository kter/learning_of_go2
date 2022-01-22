package main

import (
	"fmt"
)

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	// これはできない
	// fmt.Println(err.Message)
	// これはできる (型アサーションする)
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.Message)
	}
}
