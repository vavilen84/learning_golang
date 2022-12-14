package main

import (
	"fmt"
	"reflect"
)

/*
 ORM example: https://github.com/vavilen84/gocommerce/blob/0.0.1/server/orm/orm.go#L88
*/

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
