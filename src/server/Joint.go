package server

import (
	"log"
)

type Joint struct {
}

func initComponent() {
	api_map := make(HttpIFMap)
	http_server := HttpServer{
		api_map: HttpIFMap{},
	}

	ret := http_server.Init(api_map)
	if ret {
		http_server.openServer("", 8001)
		log.Println("Init http server OK")
	} else {
		log.Println("Init http server Failed")
	}
}

func (*Joint) Start() {
	initComponent()
}
