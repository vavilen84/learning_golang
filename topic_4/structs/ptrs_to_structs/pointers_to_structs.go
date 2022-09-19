package main

import "fmt"

/*
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that
notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
*/

type B struct {
	X int
	Y int
}

func main() {
	v := B{1, 2}
	p := &v
	p.X = 1
	fmt.Println(v)
}
