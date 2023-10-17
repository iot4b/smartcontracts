package api

import (
	"encoding/json"
	"github.com/coalalib/coalago"
	"github.com/ever-iot/node/db"
	"github.com/ever-iot/node/dsm"
	"github.com/ever-iot/node/everscale"
	"github.com/ever-iot/node/helpers"
	"github.com/ever-iot/node/locator"
	"github.com/ever-iot/node/system/config"
	"github.com/ever-iot/node/utils"
	"github.com/jinzhu/copier"
	log "github.com/ndmsystems/golog"
	"github.com/pkg/errors"
)

// метод /register принимает на вход данные устройства для регистрации в блокчейне
type registerReq struct {
	Address    string `json:"a"`            // адрес контракта если есть
	Vendor     string `json:"v"`            // адрес вендора
	Version    string `json:"ver"`          // версия прошивки
	Type       string `json:"t,omitempty"`  // название модели устройства
	VendorData string `json:"vd,omitempty"` // произволный блок данных в любом формате
}

type registerRes struct {
	Address dsm.EverAddress `json:"a,omitempty"` //ever SC address текущего Device
	Node    dsm.EverAddress `json:"n,omitempty"` //ever SC address Node, с которой девайс создал последнее соединение
	Elector dsm.EverAddress `json:"e,omitempty"` //ever SC адрес Elector'a, который обслуживает сеть нод для текущего девайса
	Vendor  dsm.EverAddress `json:"v,omitempty"` //ever SC address производителя текущего девайса

	Stat bool `json:"s,omitempty"` // нужно ли девайсу слать статистику
}

func getDeviceInfo(msg *coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {
	address := msg.GetURIQuery("a")

	accInfo, err := everscale.GetAccountInfo(address)
	if err != nil {
		log.Error(err)
		return coalago.NewResponse(coalago.NewStringPayload(err.Error()), coalago.CoapCodeInternalServerError)
	}

	log.Debugw("get device info", "address", address, "info", accInfo)
	result, err := json.Marshal(accInfo)
	if err != nil {
		log.Error(err)
		return coalago.NewResponse(coalago.NewStringPayload(err.Error()), coalago.CoapCodeInternalServerError)
	}
	return coalago.NewResponse(coalago.NewBytesPayload(result), coalago.CoapCodeContent)
}

// registerDevice - регистрируем девайс в блокчейне и устанавливаем текущую ноду как currentNode
// если девайс уже ранее подключался, то его запись смарт будет в блокчейне, тогда девайс передает
// ...?a= ... адрес своего контракта.
func registerDevice(msg *coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {
	log.Info("Register new device")

	// device contract
	//device := everscale.Device{}

	// если адрес пустой, значит девайс не знает свой адрес. деплоим контракт нового девайса
	// todo поиск контракта в блокчейне

	req := registerReq{}
	err := json.Unmarshal(msg.Payload.Bytes(), &req)
	if err != nil {
		log.Error(err, msg.Payload.String())
		return helpers.OutputMessage(coalago.CoapCodeInternalServerError, err.Error())
	}
	log.Debug(msg.Payload.String())

	// todo какие поля нода не может изменять у девайса?
	// todo нужно где-то хранить адрес ноды
	// todo на данном этапе можно организовать взаимодействие с блокчейном напрямую с девайса

	// берем из глобального пакета данные по текущей ноде
	// если девайс повторно подключается или первый раз, то некоторые данные нужно обязательно перезаписать (при каждом подключении к ноде):
	// - Node адрес текущей ноды (на которой происходит регистрация)
	// - Elector фиксированный, его может менять только вендор
	// - Vendor фиксированный. По нему происходит привязка к контракту производителя, у которого есть какие-то права над контрактом девайса
	// - Version - динамическое. Обновлять при каждом подключении
	// - Active при подключении устанавливаем в по-умолчанию true, далее за этим флажком следит /alive.
	//   todo продумать как пачкой менять значения для группы девайсов
	// - VendorData - динамичный. по-умолчанию с девайса прилетает
	//   todo но может меняться только Vendor'о?

	//copier.Copy(&device, req)

	if len(req.Address) == 0 {
		log.Debug(req)
		return helpers.OutputMessage(coalago.CoapCodeBadRequest, "contract address not provided")
	}

	// обновляем ноду на текущую
	//device.Node = node.Get().Address
	//device.Elector = node.Get().Elector
	//device.Owners = []dsm.EverAddress{device.Node}
	//device.VendorName = "VendorName"
	//device.VendorData = "VendorData"
	//
	//device.Active = true

	// deploy
	//var signerKeys *domain.KeyPair // owner keys for deployed contract
	//if len(req.Address) == 0 {
	//	log.Debug(device)
	//	signerKeys, err = device.Deploy(1000000000)
	//	if err != nil {
	//		log.Error("device.Deploy:", err)
	//		return helpers.OutputMessage(coalago.CoapCodeInternalServerError, err.Error())
	//	}
	//
	//}

	// save to db
	d := db.Device{
		Address: req.Address,
		Type:    req.Type,
	}
	if err = d.Save(); err != nil {
		return helpers.OutputMessage(coalago.CoapCodeInternalServerError, "error saving device: "+err.Error())
	}

	// после сохранения данных о девайсе в БД и блокчейн, возвращаем результат регистрации
	result := registerRes{}
	copier.Copy(&result, req)
	//copier.Copy(&result, signerKeys)

	log.Infow("New device registered: address", req.Address, "type", req.Type)

	return helpers.OutputJson(coalago.CoapCodeContent, result)
}

func locateDevice(msg *coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {
	address := msg.GetURIQuery("address")
	_, err := db.Get(address)
	if errors.Is(err, db.ErrNotFound) {
		nodeIp, err := locator.Locate(address)
		if err != nil {
			return coalago.NewResponse(coalago.NewStringPayload("Address not found: "+err.Error()), coalago.CoapCodeNotFound)
		}
		return coalago.NewResponse(coalago.NewStringPayload(nodeIp), coalago.CoapCodeContent)
	} else if err != nil {
		return coalago.NewResponse(coalago.NewStringPayload(err.Error()), coalago.CoapCodeInternalServerError)
	}
	return helpers.OutputMessage(coalago.CoapCodeContent, utils.MyIp()+":"+config.Get("coala.port"))
}
