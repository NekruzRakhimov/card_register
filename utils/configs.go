package utils

import (
	"card_register/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	AppSettings models.Settings
)

func ReadSettings() {
	fmt.Println("Starting reading settings file")
	configFile, err := os.Open("./settings.json") // поче
	if err != nil {
		log.Fatal("Couldn't open config file. Error is: ", err.Error())
	}

	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			log.Fatal("Couldn't close config file. Error is: ", err.Error())
		}
	}(configFile)

	fmt.Println("Starting decoding settings file")
	if err = json.NewDecoder(configFile).Decode(&AppSettings); err != nil {
		log.Fatal("Couldn't decode settings json file. Error is: ", err.Error())
	}

	log.Println(AppSettings)
	return
}
