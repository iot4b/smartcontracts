package api

import (
	"encoding/json"
	"errors"
	"github.com/coalalib/coalago"
	"github.com/ever-iot/node/db"
	"github.com/ever-iot/node/everscale"
	"github.com/ever-iot/node/smartmoc"
	"math/rand"
	"time"

	log "github.com/ndmsystems/golog"
)

type cmd struct {
	Address string `json:"address"`
	Cmd     string `json:"cmd"`
}

// info для коалы
func handlerInfo(i map[string]interface{}) func(message *coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {
	return func(message *coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {
		// add random delay to imitate various ping time
		r := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * r)

		info, err := json.Marshal(i)
		if err != nil {
			log.Errorw(err.Error(), "info", i)
			return coalago.NewResponse(coalago.NewStringPayload(err.Error()), coalago.CoapCodeBadRequest)
		}
		log.Debug(string(info))
		return coalago.NewResponse(coalago.NewBytesPayload(info), coalago.CoapCodeContent)
	}
}

// alive  пакеты
func alive(message *coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {
	address := message.GetURIQuery("a")
	log.Debugw("alive address", "address", address, "ip", message.Sender.String())
	if err := db.Alive(address, message.Sender.String()); err != nil {
		//не отвечаем на запрос клиенту
		return nil
	}

	return coalago.NewResponse(coalago.NewStringPayload(message.Sender.String()), coalago.CoapCodeChanged)
}

func sendCmd(message *coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {
	p := cmd{}
	json.Unmarshal(message.GetPayload(), &p)
	log.Debug("sendCmd", p)

	if len(p.Cmd) == 0 {
		log.Debug("Empty cmd", p.Address)
		return coalago.NewResponse(coalago.NewStringPayload("Empty cmd"), coalago.CoapCodeBadRequest)
	}

	device, err := db.Get(p.Address)
	if errors.Is(err, db.ErrNotFound) {
		return coalago.NewResponse(coalago.NewStringPayload("Address not found"), coalago.CoapCodeNotFound)
	} else if err != nil {
		return coalago.NewResponse(coalago.NewStringPayload(err.Error()), coalago.CoapCodeInternalServerError)
	}

	ipPort := device.IpPort

	log.Debug("address:", p.Address, "ip:", ipPort, "cmd:", p.Cmd)

	if !device.Online || len(ipPort) == 0 {
		return coalago.NewResponse(coalago.NewStringPayload("Device is offline"), coalago.CoapCodeServiceUnavailable)
	}

	message.SetProxy(message.GetSchemeString(), ipPort)
	resp, err := client.Send(message, "127.0.0.1:"+coalaPort)
	if err != nil {
		log.Error(err, ipPort)
		return coalago.NewResponse(coalago.NewStringPayload(err.Error()), coalago.CoapCodeInternalServerError)
	}

	log.Debug(string(resp.Body))
	return coalago.NewResponse(coalago.NewStringPayload(string(resp.Body)), coalago.CoapCodeContent)
}

// getEndpoints - возвращаем список зарегистрированных нод в блокчейне
func getEndpoints(_ *coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {
	var payload []byte
	elector := smartmoc.Elector{}
	nodeList := elector.GetCurrentPeriodNodes()
	payload, _ = json.Marshal(nodeList)

	return coalago.NewResponse(coalago.NewBytesPayload(payload), coalago.CoapCodeContent)
}

// getDeviceContract - получить контракт девайса
func getDeviceContract(_ *coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {

	return nil
}

func getAccountInfo(msg *coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {
	address := msg.GetURIQuery("a")

	accInfo, err := everscale.GetAccountInfo(address)
	if err != nil {
		log.Error(err)
		return coalago.NewResponse(coalago.NewStringPayload(err.Error()), coalago.CoapCodeInternalServerError)
	}

	log.Debugw("get account info", "address", address, "info", accInfo)
	result, err := json.Marshal(accInfo)
	if err != nil {
		log.Error(err)
		return coalago.NewResponse(coalago.NewStringPayload(err.Error()), coalago.CoapCodeInternalServerError)
	}
	return coalago.NewResponse(coalago.NewBytesPayload(result), coalago.CoapCodeContent)
}

func getContract(msg *coalago.CoAPMessage) *coalago.CoAPResourceHandlerResult {
	address := msg.GetURIQuery("a")

	log.Debugw("get contract", "address", address)
	return nil
}
