package config

import (
	"log"
	"os"
)

func LoadConfig(x string) (string, int) {
	file_hdl, err := os.Open(x)
	if file_hdl != nil {
		log.Println("There is error while reading file: %", err)
	} else {
		defer file_hdl.Close()
	}
	return "", 0
}
