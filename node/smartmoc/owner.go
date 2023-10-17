package smartmoc

import "github.com/ever-iot/node/dsm"

type Owner struct {
	Account string          `json:"account"`           // blockchain account address
	Elector dsm.EverAddress `json:"elector,omitempty"` //address
}
