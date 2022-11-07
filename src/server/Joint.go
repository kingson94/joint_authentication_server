package server

import (
	"log"
)

type Joint struct {
}

func initComponent() {
	api_map := make(HttpIFMap)
	server_body := HttpServer{
		api_map: HttpIFMap{},
	}

	ret := server_body.Init(api_map)
	if ret {
		server_body.openServer("", 8001)
		log.Println("Init http server OK")
	} else {
		log.Println("Init http server Failed")
	}
}

func (*Joint) Start() {
	initComponent()
}
