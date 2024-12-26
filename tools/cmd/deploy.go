package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/markgenuine/ever-client-go/domain"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"smartcontracts/everscale"
	"smartcontracts/shared/config"
)

// deployCmd deploys contract
var deployCmd = &cobra.Command{
	Use:   "deploy {name}",
	Short: "Deploy {name} smartcontract",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("not enough arguments")
		}

		name := args[0]

		abi, tvc, err := everscale.ReadContract(name)
		if err != nil {
			return fmt.Errorf("everscale.ReadContract: %w", err)
		}

		keyPair, err := everscale.Ever.Crypto.GenerateRandomSignKeys()
		if err != nil {
			return fmt.Errorf("everscale.Ever.Crypto.GenerateRandomSignKeys: %w", err)
		}
		fmt.Println("keyPair:", keyPair)

		// vendor contract params
		var constructorParams = struct {
			Elector     everscale.EverAddress `json:"elector"`
			VendorName  string                `json:"vendorName"`
			ProfitShare int                   `json:"profitShare"`
			ContactInfo string                `json:"contactInfo"`
		}{
			Elector:     everscale.EverAddress(config.Get("everscale.elector")),
			VendorName:  "Vendor",
			ProfitShare: 50,
			ContactInfo: "",
		}

		deployParams := &domain.ParamsOfEncodeMessage{
			Abi:    abi,
			Signer: domain.NewSigner(domain.SignerNone{}),
			DeploySet: &domain.DeploySet{
				Tvc:           base64.StdEncoding.EncodeToString(tvc),
				InitialPubKey: keyPair.Public,
			},
			CallSet: &domain.CallSet{
				FunctionName: "constructor",
				Input:        constructorParams,
			},
		}

		msg, err := everscale.Ever.Abi.EncodeMessage(deployParams)
		if err != nil {
			return fmt.Errorf("everscale.Ever.Abi.EncodeMessage: %w", err)
		}

		if err = everscale.Giver.SendTo(msg.Address, 10_000_000_000); err != nil {
			return fmt.Errorf("Giver.SendTo(%s): %w", msg.Address, err)
		}

		params := &domain.ParamsOfProcessMessage{
			MessageEncodeParams: deployParams,
			SendEvents:          false,
		}

		fmt.Println("Deploying contract...", msg.Address)
		resp, err := everscale.Ever.Processing.ProcessMessage(params, nil)
		if err != nil {
			return fmt.Errorf("everscale.Ever.Processing.ProcessMessage: %w", err)
		}

		fmt.Println(string(resp.Decoded.Output))

		return nil
	},
}

func init() {
	RootCmd.AddCommand(deployCmd)
}
