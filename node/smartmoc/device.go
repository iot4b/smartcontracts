package smartmoc

import (
	"github.com/ever-iot/node/dsm"
)

type Device struct {
	Address dsm.EverAddress   `json:"address,omitempty"` //ever address of smart contract
	Node    dsm.EverAddress   `json:"node,omitempty"`    //address, изменяется
	Elector dsm.EverAddress   `json:"elector,omitempty"` //address
	Vendor  dsm.EverAddress   `json:"vendor,omitempty"`  //address
	Owners  []dsm.EverAddress `json:"owners,omitempty"`  //address

	//изменяемые параметры
	Active  bool   `json:"active,omitempty"`  //активен ли девайс в iot, изменяется
	Lock    bool   `json:"lock,omitempty"`    //статус залочен ли, изменяется
	Stat    bool   `json:"stat,omitempty"`    //слать ли статистику, изменяется
	Version string `json:"version,omitempty"` //версия прошивки, изменяется

	PublicKey  string `json:"publicKey,omitempty"`
	Type       string `json:"type,omitempty"`       //тип девайса
	VendorName string `json:"vendorName,omitempty"` //название вендора
	VendorData string `json:"vendorData,omitempty"` //происзолный блок данных вендора в зашифрованном виде
}

// GetNodeAddress from device contract data
func (d *Device) GetNodeAddress() dsm.EverAddress {
	return ""
}

// SetNodeAddress to device contract data
func (d *Device) SetNodeAddress(node dsm.EverAddress) error {
	d.Node = node
	return nil
}

// Pay - оплата периода обслуживания
// метод может вызвать только нода прописанная в смарт
func (d *Device) Pay() bool {
	//получаем с электора стоимость текущего раунда
	//проверяем что нода пытающаяся списать деньги прописана как управляющая
	//делаем транзакцию
	return true
}

// CheckLock - проверяет девайс в блоке или нет
func (d *Device) CheckLock() bool {
	return d.Lock
}

// SetLock - блокирует работу девайса
func (d *Device) SetLock() {
	d.Lock = true
}

// Unlock - разблокирует работу девайса
func (d *Device) Unlock() {
	d.Lock = false
}
