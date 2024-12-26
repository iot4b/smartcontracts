package everscale

import (
	"fmt"
)

type giver struct {
	Address EverAddress // blockchain account address
	Public  string
	Secret  string
}

var Giver giver

// SendTo transfer a [value] of test nanotokens from giver.Address to [account]
func (g giver) SendTo(account string, value int) (err error) {
	abi, err := getAbiFromFile("../build/Giver.abi.json")
	if err != nil {
		return fmt.Errorf("getAbiFromFile(../build/Giver.abi.json): %w", err)
	}

	input := sendTransaction{
		Dest:   account,
		Value:  value,
		Bounce: false, // for deploy
	}
	_, err = processMessage(
		abi,
		string(g.Address),
		"sendTransaction",
		input,
		NewSigner(g.Public, g.Secret),
	)
	return
}
