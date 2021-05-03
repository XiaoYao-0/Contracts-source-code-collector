package conf

import (
	"encoding/json"
	"errors"
	"os"
)

type conf struct {
	ContractNumber int    `json:"contract_number"`
	Proxy          string `json:"proxy"`
	ApiKey         string `json:"api_key"`
	StorageDir     string `json:"storage_dir"`
}

func getConf() (conf, error) {
	var conf conf
	file, err := os.Open("conf.json")
	if err != nil {
		return conf, errors.New("conf.json not found")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return conf, errors.New("unresolved conf.json")
	}
	return conf, nil
}

func ContractNumber() (int, error) {
	var num int
	conf, err := getConf()
	if err != nil {
		return num, err
	}
	return conf.ContractNumber, nil
}

func SetContractNumber(num int) error {
	var conf conf
	file, err := os.Open("conf.json")

	if err != nil {
		return errors.New("conf.json not found")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return errors.New("unresolved conf.json")
	}
	file.Close()
	os.Truncate("conf.json", 0)
	file, err = os.OpenFile("conf.json", os.O_APPEND, 0644)
	conf.ContractNumber = num
	encoder := json.NewEncoder(file)
	err = encoder.Encode(&conf)
	if err != nil {
		return errors.New("write conf.json failure")
	}
	return nil
}

func Proxy() (string, error) {
	var proxy string
	conf, err := getConf()
	if err != nil {
		return proxy, err
	}
	return conf.Proxy, nil
}

func ApiKey() (string, error) {
	var apiKey string
	conf, err := getConf()
	if err != nil {
		return apiKey, err
	}
	if conf.ApiKey == "" {
		return apiKey, errors.New("api key empty")
	}
	return conf.ApiKey, nil
}

func StorageDir() (string, error) {
	var storageDir string
	conf, err := getConf()
	if err != nil {
		return storageDir, err
	}
	return conf.StorageDir, nil
}
