package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
}

func main() {
	u := User{}
	b, err := json.Marshal(u)
	fmt.Printf("Err %v \r\n", err)
	fmt.Printf("Bytes %v \r\n", b)
	fmt.Printf("Marshal %v \r\n", string(b))

	u = User{
		FirstName: "John",
		LastName:  "Dou",
	}
	b, _ = json.Marshal(u)
	fmt.Printf("Marshal %v \r\n", string(b))

	u = User{}
	json.Unmarshal(b, &u)
	fmt.Printf("Unmarshal %+v \r\n", u)

	u = User{}
	b = []byte(`{"first_name":"John","last_name":"Dou"} `)
	json.Unmarshal(b, &u)
	fmt.Printf("Unmarshal %+v \r\n", u)
}
