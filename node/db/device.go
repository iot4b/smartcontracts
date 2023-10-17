package db

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/ndmsystems/golog"
	bolt "go.etcd.io/bbolt"
)

type Device struct { // json короткие имена, потому что это сохранит гигабайты на диске
	Address    string `json:"a"` // адрес контракта девайса
	Type       string `json:"t"` // тип девайса
	IpPort     string `json:"i"` // текущий адрес девайса. внешний ip:port для отправки через coala по UDP
	Online     bool   `json:"o"` // онлайн или нет. если девайс не шлет alive в течении device.aliveTimeout, то становится оффлайн
	LastUpdate int64  `json:"l"` // дата обновления данных
}

// Alive - регистрируем alive
func Alive(deviceAddress, ipport string) error {
	log.Debug(deviceAddress, ipport)

	update := false //флаг, по которому принимаем решение, надо ли апдейтить запись в общей базе данных

	d, err := Get(deviceAddress)
	if err != nil {
		// device is not registered
		log.Error("device not registered", err, deviceAddress)
		return err
	}

	if d.IpPort != ipport {

		//детектим изменения ip/port.
		if d.IpPort != "" {
			ipData := strings.Split(ipport, ":")
			if len(ipData) != 2 {
				log.Debug("ipport change metric fault current:", d.IpPort, "new:", ipport)
				ipData = []string{"", ""}
			}

			if !strings.Contains(d.IpPort, ipData[0]) {
				IpPortChange.WithLabelValues("ip").Inc()
			}

			if !strings.Contains(d.IpPort, ipData[1]) {
				IpPortChange.WithLabelValues("port").Inc()
			}
		}

		d.IpPort = ipport
		update = true
	}

	if d.Online == false {
		d.Online = true
		update = true
	}

	// update device in online devices list
	onlineList.SetDefault(d.Address, d)

	if update || !d.Exists() {
		return d.Save()
	}

	return nil
}

// Save device data to DB
func (d Device) Save() error {
	if d.Address == "" {
		return errors.New("empty address")
	}

	isNew := false
	if d.LastUpdate == 0 {
		isNew = true
	}

	d.LastUpdate = time.Now().Unix()

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(devicesBucket))
		data, err := json.Marshal(d)
		if err != nil {
			return err
		}
		err = b.Put([]byte(d.Address), data)
		if err != nil {
			log.Errorw(err.Error(), "d", d, "data", string(data))
		}
		return err
	})
	if err != nil {
		return err
	}

	// metrics
	WritesCount.Inc()
	if isNew {
		DevicesNum.Inc()
	}

	return nil
}

// Exists check if device exists in DB
func (d Device) Exists() (exists bool) {
	log.Debug(d)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(devicesBucket))
		data := b.Get([]byte(d.Address))
		exists = data != nil
		return nil
	})
	return exists
}

// devicesCount number of devices in DB
func devicesCount() (count int) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(devicesBucket))
		count = b.Stats().KeyN
		return nil
	})
	return
}
