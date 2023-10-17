package dsm

// EverAddress - формат адреса в everscale,
// для приведения к единому формату и валидация
type EverAddress string

type Node struct {
	Address EverAddress `json:"address,omitempty"`
	Elector EverAddress `json:"elector,omitempty"`
	Owner   EverAddress `json:"owner,omitempty"`

	IpPort string `json:"ipPort,omitempty"`
}

type Device struct {
	Address EverAddress `json:"address,omitempty"` //ever address of smart contract
	Active  bool        `json:"active,omitempty"`  //активен ли девайс в iot, изменяется
	Lock    bool        `json:"lock,omitempty"`    //статус залочен ли, изменяется
	Stat    bool        `json:"stat,omitempty"`    //слать ли статистику, изменяется
	Version string      `json:"version,omitempty"` //версия прошивки, изменяется
	Type    string      `json:"type,omitempty"`    //тип девайса
	Node    string      `json:"node"`              // node ip address
}
