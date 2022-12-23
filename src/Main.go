package main

import (
	"log"
	"os"
	"os/signal"
	"src/server"
	"syscall"
	"time"
)

func signalHandle(signal os.Signal) {
	if signal == syscall.SIGTERM {
		log.Println("Receive termination signal:", signal.String())
		// Handle signal here
		os.Exit(0)
	} else if signal == syscall.SIGINT {
		log.Println("Receive closing signal:", signal.String())
		// Handle signal here
		os.Exit(0)
	} else {
		// log.Println("Ignoring signal:", signal)
	}
}

func main() {
	server := server.Joint{}
	if len(os.Args) > 1 {
		server.ConfigFilePath = os.Args[1]
	}
	server.Start()

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan)

	go func() {
		for {
			sign := <-signal_chan
			signalHandle(sign)
		}
	}()

	for {
		time.Sleep(10 * time.Millisecond)
	}
}
