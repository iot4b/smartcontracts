package cmd

import "smartcontracts/everscale"

type (
	Elector struct {
		DefaultNodes []everscale.EverAddress `json:"defaultNodes"`
	}

	Vendor struct {
		Elector     everscale.EverAddress `json:"elector"`
		VendorName  string                `json:"vendorName"`
		ProfitShare int                   `json:"profitShare"`
		ContactInfo string                `json:"contactInfo"`
	}
)
