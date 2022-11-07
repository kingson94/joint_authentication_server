package core

import (
	"log"
	"src/model"
)

type Worker struct {
	id        int
	id_string string
	channel   chan (string)
}

func Process(pTask *model.Task) {
	p := model.Task{
		Data: []byte("data"),
	}
	log.Println("Data ", &pTask.Data, p)
}
