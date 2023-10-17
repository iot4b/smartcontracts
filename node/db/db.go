package db

import (
	"encoding/json"
	"errors"
	"github.com/patrickmn/go-cache"
	"time"

	"github.com/ndmsystems/golog"
	bolt "go.etcd.io/bbolt"
)

const devicesBucket = "devices"

var (
	ErrNotFound = errors.New("not found")

	db         *bolt.DB     // bolt db instance
	onlineList *cache.Cache // memory storage for online devices
)

func Init(boltDbPath string, aliveTimeout time.Duration) {
	log.Info("Init DB")

	var err error
	// open bolt db
	db, err = bolt.Open(boltDbPath, 0600, &bolt.Options{Timeout: time.Minute})
	if err != nil {
		log.Fatal("Cannot open BoltDB:", err)
	}

	// ensure devices bucket exists
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(devicesBucket))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	//инитим мемори базу, deviceAddress будет жить в ней 1 минуту согласно конфигу.
	onlineList = cache.New(aliveTimeout, time.Second*60)

	//шлем событие из базы когда девайс становится офлайн
	onlineList.OnEvicted(func(deviceAddress string, data interface{}) {
		log.Debug("deviceAddress expired", deviceAddress)
		d, _ := Get(deviceAddress)
		d.Online = false
		d.IpPort = ""
		d.Save()
	})

	// сохраняем в метрики количество устройств
	DevicesNum.Set(float64(devicesCount()))
}

// Close DB connection when finished
func Close() {
	if err := db.Close(); err != nil {
		log.Error("db.Close:", err)
	}
}

// Get device data by deviceAddress
func Get(deviceAddress string) (d Device, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(devicesBucket))
		data := b.Get([]byte(deviceAddress))
		if data == nil {
			log.Debug(ErrNotFound, deviceAddress)
			return ErrNotFound
		}
		err = json.Unmarshal(data, &d)
		if err != nil {
			log.Error(err)
		}
		return err
	})
	return
}
