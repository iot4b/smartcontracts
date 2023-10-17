package locator

import (
	"encoding/json"
	"github.com/ever-iot/node/dsm"
	"github.com/ever-iot/node/everscale"
	log "github.com/ndmsystems/golog"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"time"
)

type (
	device struct {
		Address dsm.EverAddress `json:"address"`
		Node    string          `json:"node"`
	}
	response struct {
		Value string `json:"value0"`
	}
)

var devices = cache.New(time.Hour, time.Minute*10)

// Locate ip:port of the device node
func Locate(deviceContractAddress string) (string, error) {
	log.Debug("Locate:", deviceContractAddress)
	if len(deviceContractAddress) == 0 {
		return "", errors.New("contract address not provided")
	}
	// check for device data in cache
	d, ok := devices.Get(deviceContractAddress)
	if ok {
		return d.(device).Node, nil
	}

	// find node contract address
	deviceRes, err := everscale.Execute("device", dsm.EverAddress(deviceContractAddress), "getNode", nil)
	if err != nil {
		log.Error(errors.Wrap(err, "everscale.Execute(getNode)"))
		return "", err
	}
	dr := response{}
	if err = json.Unmarshal(deviceRes, &dr); err != nil {
		return "", err
	}

	// get node ip:port from node contract
	nodeRes, err := everscale.Execute("node", dsm.EverAddress(dr.Value), "getIpPort", nil)
	if err != nil {
		log.Error(errors.Wrap(err, "everscale.Execute(getIpPort)"))
		return "", err
	}
	nr := response{}
	if err = json.Unmarshal(nodeRes, &nr); err != nil {
		return "", err
	}

	// save to cache
	devices.SetDefault(deviceContractAddress, device{
		Address: dsm.EverAddress(deviceContractAddress),
		Node:    nr.Value,
	})

	return nr.Value, nil
}
