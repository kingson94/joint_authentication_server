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
	api_map   HttpIFMap
	serve_mux *http.ServeMux
}

func (http_server *HttpServer) Init(api_map HttpIFMap) bool {
	http_server.api_map = api_map
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

func (http_server *HttpServer) openServer(host string, port int) {
	serve_mux := http.NewServeMux()

	for ifs, cb := range http_server.api_map {
		serve_mux.HandleFunc(ifs, cb)
	}

	go func() {
		err := http.ListenAndServe(host+":"+strconv.Itoa(port), serve_mux)
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
