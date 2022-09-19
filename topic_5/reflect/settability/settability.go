package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x) // important - need to pass pointer
	fmt.Printf("%v \r\n", v.Elem().CanSet())
	v.Elem().SetFloat(7.1)
	fmt.Printf("%v \r\n", v.Elem().Interface())
	// or
	fmt.Printf("%v \r\n", reflect.Indirect(v).Float())
}
