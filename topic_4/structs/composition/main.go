package main

import "fmt"

type ShippingProvider struct {
	apiKey string
}

func (s *ShippingProvider) SetApiKey(apiKey string) {
	s.apiKey = apiKey
}

func (s *ShippingProvider) GetApiKey() string {
	return s.apiKey
}

type USPS struct {
	ShippingProvider
	USPSSpecificData string
}

func main() {
	usps := USPS{
		USPSSpecificData: "USPS_specific_data",
	}
	usps.SetApiKey("sd0987sdf0897sdv0987dsvsd9v0")
	fmt.Println(usps.GetApiKey())
}
