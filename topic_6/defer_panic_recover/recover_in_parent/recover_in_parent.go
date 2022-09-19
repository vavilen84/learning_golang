package main

import "fmt"

func e() {
	panic(100)
}

func d() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in d", r)
		}
	}()
	e()
}

func main() {
	d()
	fmt.Println("main")
}
