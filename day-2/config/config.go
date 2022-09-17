package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var configEnv map[string]string

func init() {
	content, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Println("config.json file not found")
	} else {
		err = json.Unmarshal(content, &configEnv)
		if err != nil {
			log.Println("invalid config.json file")
		}
	}
}

func GetValue(key string) string {
	value, ok := configEnv[key]
	if !ok {
		return ""
	}
	return value

}
