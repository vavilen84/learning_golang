package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("x value: ", reflect.ValueOf(x))
	fmt.Printf("Value %+v \r\n", reflect.ValueOf(x))
	fmt.Printf("Value Type %T \r\n", reflect.ValueOf(x))
	fmt.Println("x Float value: ", reflect.ValueOf(x).Float())
	fmt.Printf("Value Float Type %T \r\n", reflect.ValueOf(x).Float())
	fmt.Println("x value type: ", reflect.ValueOf(x).Type())
	fmt.Printf("Value Type obj %T \r\n", reflect.ValueOf(x).Type())
	fmt.Println("x value kind: ", reflect.ValueOf(x).Kind())

	y, ok := reflect.ValueOf(x).Interface().(float64)
	fmt.Printf("%v %T %v \r\n", y, y, ok)
}
