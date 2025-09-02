package main

import (
	"errors"
	"fmt"
)

func customerr() {
	if err := doSomething(); err != nil {
		fmt.Println(err)
		return
	}
}

type customError struct {
	code    int32
	message string
	err     error
}

// error returns the error message
// implementing error interface
func (m *customError) Error() string {
	return fmt.Sprintf("Error %d: %s %v \n", m.code, m.message, m.err)
}

//function that return cistomError

//func doSomething()error{
//return &customError{
//	code :500,
//	message: "Something went wrong",
//}
//}

func doSomething() error {
	if err := doSomethingElse(); err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			err:     err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("Internal Error")
}
