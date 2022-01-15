package helper

import (
	"encoding/json"
	config "getir-case/config"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadConfig(env string) *config.AppConfig {

	f, err := os.Open("settings.json")
	if err != nil {
		log.Fatalf("An error occurred while opening settings.json. Error : %s", err.Error())
	}
	defer f.Close()

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		log.Fatalf("An error occurred while reading settings.json. Error : %s", err.Error())
	}
	settings := make(map[string]*config.AppConfig)
	err = json.Unmarshal(data, &settings)
	if err != nil {
		log.Fatalf("An error occurred while parsing settings.json. Error %s", err.Error())
	}

	cfg, exist := settings[env]
	if !exist {
		log.Fatalf("Invalid envrionment %s", env)
	}

	cfg.Port = readPort()

	return cfg
}

func readPort() string {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return port
}

//Returns GO_ENV value of your OS. Default value is "local"
func ReadEnv() string {

	env := os.Getenv("GO_ENV")

	if strings.TrimSpace(env) == "" {
		env = "local"
	}
	return env
}
