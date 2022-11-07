package server

import (
	"io"
	"log"
	"net/http"
	"src/protogen"
	"strconv"

	"google.golang.org/protobuf/encoding/protojson"
)

type HttpIFMap map[string]func(http.ResponseWriter, *http.Request)
type HttpServer struct {
	api_map HttpIFMap
}

func (server_body *HttpServer) Init(api_map HttpIFMap) bool {
	server_body.api_map = api_map
	api_map["/"] = Welcome
	api_map["/register/new_account"] = RegisterNewAccount

	// Init server setting
	name := "Son"
	phone := "098398291"
	user := "sontv"
	pass := "s123"
	p := protogen.Profile{
		Name:     &name,
		Tel:      &phone,
		Username: &user,
		Password: &pass,
	}

	msg_byte, _ := protojson.Marshal(&p)
	log.Println(msg_byte)
	log.Println(string(msg_byte))

	return true
}

func (server_body *HttpServer) openServer(host string, port int) {
	http_server := http.NewServeMux()

	for ifs, cb := range server_body.api_map {
		http_server.HandleFunc(ifs, cb)
	}

	go func() {
		err := http.ListenAndServe(host+":"+strconv.Itoa(port), http_server)
		log.Println("Start http server failed: ", err)
	}()
}

func Welcome(res_writer http.ResponseWriter, request *http.Request) {
	io.WriteString(res_writer, "Welcome to authentication server")
}

func RegisterNewAccount(res_writer http.ResponseWriter, request *http.Request) {
	io.WriteString(res_writer, "Register succ")
	defer request.Body.Close()
	b, _ := io.ReadAll(request.Body)
	log.Println(string(b))
}
