package utils

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"os"
)

var (
	ErrUnmarshal = errors.New("file unmarshal error")
)

func ReadJSONFile(path string, to any) error {
	fileData, err := ReadFile(path)
	if err != nil {
		return errors.Wrapf(err, "ReadFile(%s)", path)
	}
	err = json.Unmarshal(fileData, to)
	if err != nil {
		return errors.Wrapf(ErrUnmarshal, "json.Unmarshal(%s, to): %s", path, err.Error())
	}
	return nil
}

func SaveFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "os.Open(%s)", path)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Wrapf(err, "io.ReadAll(%s)", path)
	}
	return bytes, nil
}
