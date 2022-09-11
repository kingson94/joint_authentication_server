package http_server

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func OpenServer(host string, port int) {
	http_server := http.NewServeMux()
	http_server.HandleFunc("/", Welcome)
	http_server.HandleFunc("/register/v1", RegisterV1)

	go func() {
		err := http.ListenAndServe(host+":"+strconv.Itoa(port), http_server)
		log.Println("Start http server failed: ", err)
	}()

}

func Welcome(res_writer http.ResponseWriter, request *http.Request) {
	io.WriteString(res_writer, "Welcome to authentication server")
}

func RegisterV1(res_writer http.ResponseWriter, request *http.Request) {
	io.WriteString(res_writer, "Register succ")
}
