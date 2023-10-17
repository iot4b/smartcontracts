package smartmoc

type Elector struct {
	Account string `json:"account"` // blockchain account address
}

type node struct {
	IpPort  string `json:"ipPort"`  // network address
	Account string `json:"account"` // blockchain account address
}

var nodes = []node{
	{"157.245.57.218:5683", "0:b6f0dd040d1fa3f79871418040e7bb42c02dd674d32cc29fbc82f97dc119a212"},
	{"188.166.71.64:5683", "0:e413df4ff5e475d6dfd78db61ea635d8eefb18753f3152c51b2929edd684ec57"},
	{"192.81.216.57:5683", "0:c6bd6af3a4104c9cd3867c79cff5f62410cc95ca94e5e235c8303f9213f19582"},
	{"240.0.0.0:65535", "0:0000000000000000000000000000000000000000000000000000000000000000"}, // non-existent node
}

func (e Elector) GetCurrentPeriodNodes() []node {
	return nodes
}

func (e Elector) GetNextPeriodNodes() []node {
	return nodes
}

func (e Elector) RegisterForElection(node string) {
}

func (e Elector) ReportNodeFail(node string) {
}

func (e Elector) GetCurrentPeriodCost() float64 {
	return 0.1
}

func (e Elector) GetNextPeriodCost() float64 {
	return 0.1
}

func (e Elector) Election() {

}
