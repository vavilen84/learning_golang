package main

import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func runWithBuiltinError() error {
	return errors.New("error!")
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	err := runWithBuiltinError()
	if err != nil {
		fmt.Println(err)
	}
}
