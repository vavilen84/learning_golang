package main

import "fmt"

type Authorization interface {
	GetIdentity() []byte
}

type IdentityProvider struct{}

func (i *IdentityProvider) GetIdentity() []byte {
	return []byte(`{user:"Bob"}`)
}

type AuthProvider struct {
	Authorization
}

func main() {
	ip := IdentityProvider{}
	ap := AuthProvider{
		Authorization: &ip,
	}
	fmt.Println(string(ap.GetIdentity()))
}
