package everscale

import "github.com/ever-iot/node/dsm"

type AccountInfo struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
	Boc     string `json:"boc"`
}

// request for sendTransaction method of giver contract
type sendTransactionReq struct {
	Dest    string `json:"dest"`              // dest address
	Value   int    `json:"value"`             // amount in nano EVER
	Bounce  bool   `json:"bounce"`            // false for contract deploy
	Flags   int    `json:"flags,omitempty"`   // ???
	Payload string `json:"payload,omitempty"` // ???
}

// request for setNode method of device contract
type setNodeReq struct {
	Address dsm.EverAddress `json:"value"`
}

// response for getNode method of device contract
type getNodeRes struct {
	Address dsm.EverAddress `json:"value0"`
}

// queryResult used by ever.Net.Query() requests
type queryResult struct {
	Data struct {
		Blockchain struct {
			Account struct {
				Info AccountInfo `json:"info"`
			} `json:"account"`
		} `json:"blockchain"`
	} `json:"data"`
}
