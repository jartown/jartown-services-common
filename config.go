package common

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var Conf struct {
	Env   string `yaml:"env"`
	Debug bool   `yaml:"debug"`
	JWT   struct {
		Key     string `yaml:"key"`
		Timeout string `yaml:"timeout"`
	} `yaml:"jwt"`
}

func configLoad(filename string, conf interface{}) error {
	confData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(confData, &Conf) // common config
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(confData, conf) // app-specific config
	if err != nil {
		return err
	}

	return nil
}

func ConfigInit(conf interface{}) error {
	confFile := ""
	if len(os.Args) > 1 {
		confFile = os.Args[1]
	}
	if confFile == "" {
		confFile = os.Getenv("JARTOWN_CONFIG_FILE")
	}
	if confFile == "" {
		confFile = "../res/config.yml"
	}

	return configLoad(confFile, conf)
}
