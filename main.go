package main

import (
	"encoding/json"
	"io/ioutil"
)

const ConfigFileName = "config.json"

func Main() {
	file, err := ioutil.ReadFile(ConfigFileName)
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)

    RunBot(config)
}