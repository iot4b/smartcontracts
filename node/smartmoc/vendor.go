package smartmoc

import "github.com/ever-iot/node/dsm"

type Vendor struct {
	Account string          `json:"account"`           // blockchain account address
	Elector dsm.EverAddress `json:"elector,omitempty"` //address
}

func (v Vendor) GetName() string {
	return "ChshuJohDi and sons"
}
