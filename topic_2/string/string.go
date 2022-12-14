package main

import "fmt"

/*
 	strings in Go are read-only slices of bytes
 	a := "abc"
	a {
		len: 3,
		ptr: &0xc000014250 ptr to array
	}
*/

func main() {
	a := "abcd"
	fmt.Printf("%v \r\n", a)

	// we can re slice string
	a = a[1:3]
	fmt.Printf("%v \r\n", a)

	// correct way to iterate a string
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	aBytes := []byte(a)
	aRunes := []rune(a)
	fmt.Printf("aBytes %v \r\n", aBytes)
	fmt.Printf("aRunes %v \r\n", aRunes)

	nihongoBytes := []byte(nihongo)
	nihongoRunes := []rune(nihongo)
	fmt.Printf("nihongoBytes %v \r\n", nihongoBytes)
	fmt.Printf("nihongoRunes %v \r\n", nihongoRunes)
}
