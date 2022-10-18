package main

import "fmt"

type B struct {
	Data string
}

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
	B
	ShippingProvider
	USPSSpecificData string
}

func main() {
	usps := USPS{
		USPSSpecificData: "USPS_specific_data",
	}
	usps.SetApiKey("sd0987sdf0897sdv0987dsvsd9v0")
	fmt.Println(usps.GetApiKey())
	usps.Data = "string"

}
