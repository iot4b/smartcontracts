package node

import (
	"encoding/json"
	"github.com/ever-iot/node/dsm"
	"github.com/ever-iot/node/utils"
	"github.com/jinzhu/copier"
	log "github.com/ndmsystems/golog"
	"github.com/pkg/errors"
	"os"
)

var currentNode *dsm.Node

// todo взаимодействие с Elector'ом

func Init(localPath, coalaPort string) {
	log.Debug("Init current node data")
	currentNode = new(dsm.Node)

	localNode, err := ReadNodeFromLocal(localPath)
	// приошибке парсинга файла, либо при его отсутствии, генерим
	if err != nil && (errors.Is(err, utils.ErrUnmarshal) || errors.Is(err, os.ErrNotExist)) {
		log.Debugf("Generate mock node file with random addresses. reason: %s", err.Error())

		// todo заменить mock данные на реальные адреса в блокчейне
		// локальный файл не найден, инициируем пустой mock
		node := dsm.Node{
			Address: utils.GenerateRandomAddress(),
			Elector: utils.GenerateRandomAddress(),
			Owner:   utils.GenerateRandomAddress(),
			IpPort:  utils.MyIp() + ":" + coalaPort,
		}
		log.Debugf("Mock node: %+v", node)

		err = SaveNodeToLocal(localPath, node)
		if err != nil {
			log.Fatal(err)
		}
		Set(node)
		log.Debug("CurrentNode init from empty obj")
		return
	}

	// если все ок прочиталось из файла, то уст. текущую ноду из дампа
	Set(localNode)
	log.Debug("CurrentNode set from local dump file")
}

func Set(node dsm.Node) {
	copier.Copy(currentNode, node)
	log.Debugf("Set currentNode %+v", *currentNode)
}

func Get() *dsm.Node {
	return currentNode
}

// SaveNodeToLocal - сохраняем ноду локально
func SaveNodeToLocal(path string, n dsm.Node) error {
	data, err := json.Marshal(n)
	if err != nil {
		return errors.Wrap(err, "json.Marshal(node)")
	}
	err = utils.SaveFile(path, data)
	if err != nil {
		return errors.Wrapf(err, "utils.SaveFile(%s, data)", path)
	}
	return nil
}

// ReadNodeFromLocal - читаем ноду из локального дампа контракта
func ReadNodeFromLocal(path string) (n dsm.Node, err error) {
	err = utils.ReadJSONFile(path, &n)
	if err != nil {
		return dsm.Node{}, err
	}
	log.Debugf("%+v", n)
	return n, err
}
