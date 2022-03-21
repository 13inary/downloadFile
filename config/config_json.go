package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

// InitConfig  set globle Conf by read file
// cfg
// err error
func InitConfig() (err error) {
	confPath := ""
	flag.StringVar(&confPath, "c", "./etc/config.json", "defualt : ./etc/config.json")
	flag.Parse()

	MLock.Lock()
	ptmpConf, err := loadConfig(confPath)
	if err != nil {
		return err
	}
	Conf = *ptmpConf
	MLock.Unlock()

	return nil
}

// according path of file to get pointer of configuration struct
func loadConfig(confPath string) (conf *Config, err error) {
	// get inputs by read file
	inputs, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("read config file failed: %s", err.Error())
	}
	// transform inputs to struct
	conf = new(Config)
	err = json.Unmarshal(inputs, conf)
	if err != nil {
		return nil, fmt.Errorf("parse json failed: %s", err.Error())
	}
	return conf, nil
}
