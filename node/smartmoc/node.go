package smartmoc

import "github.com/ever-iot/node/dsm"

type Node struct {
	Account string          `json:"account"`           // blockchain account address
	Elector dsm.EverAddress `json:"elector,omitempty"` //address
	IpPort  string          `json:"ipPort"`            // network address
}

func (d *Node) GetContact() string {
	return "owneer@node.com, Alex, +1 222 333 4567"
}
