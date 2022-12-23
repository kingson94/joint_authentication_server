package server

import (
	"log"
	config "src/conf"
)

type Joint struct {
	ConfigFilePath string ""
}

func initComponent(config_path string) {
	api_map := make(HttpIFMap)
	http_server := HttpServer{
		api_map: HttpIFMap{},
	}

	config := config.Config{}
	err := config.Init(config_path)
	err |= http_server.Init(api_map)
	if err == 0 {
		server_cfg := config.ServerConfig["server"].(map[string]interface{})
		http_server.openServer("", int(server_cfg["port"].(float64)))
		log.Println("Init http server OK")
	} else {
		log.Println("Init http server Failed")
	}
}

func (c *Joint) Start() {
	initComponent(c.ConfigFilePath)
}
