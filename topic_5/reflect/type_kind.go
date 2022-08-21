package main

import (
	"fmt"
	"reflect"
)

type A struct{}
type B string

func main() {
	var x float64 = 3.4
	fmt.Printf("x type: %v %T \r\n", reflect.TypeOf(x), reflect.TypeOf(x))
	fmt.Printf("x kind: %v %T \r\n", reflect.TypeOf(x).Kind(), reflect.TypeOf(x).Kind())
	fmt.Println("x kind equal float64: ", reflect.TypeOf(x).Kind() == reflect.Float64)

	a := A{}
	fmt.Printf("a type: %v %T \r\n", reflect.TypeOf(a), reflect.TypeOf(a))
	fmt.Printf("a kind: %v %T \r\n", reflect.TypeOf(a).Kind(), reflect.TypeOf(a).Kind())
	fmt.Println("a kind equal struct: ", reflect.TypeOf(a).Kind() == reflect.Float64)

	var b B
	fmt.Printf("b type: %v %T \r\n", reflect.TypeOf(b), reflect.TypeOf(b))
	fmt.Printf("b kind: %v %T \r\n", reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Kind())
	fmt.Println("b kind equal string: ", reflect.TypeOf(b).Kind() == reflect.String)

	var c interface{}
	c = 2.3
	fmt.Printf("c type: %v %T \r\n", reflect.TypeOf(c), reflect.TypeOf(c))
	fmt.Printf("c kind: %v %T \r\n", reflect.TypeOf(c).Kind(), reflect.TypeOf(c).Kind())
	fmt.Println("c kind equal float64: ", reflect.TypeOf(c).Kind() == reflect.Float64)
}
