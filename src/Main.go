package main

import (
	http_server "src/server"
)

func main() {
	http_server.OpenServer("", 8001)
	for {
		if false {
			break
		}
	}
}
