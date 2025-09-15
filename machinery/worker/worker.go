package worker

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	Server "github.com/T-Luxshan/distributed-queue-evaluation/machinery/server"
	"log"
)

func Worker() error {
	consumerTag := "ride_worker"

	server, err := Server.StartServer()
	if err != nil {
		return err
	}

	worker := server.NewWorker(consumerTag, 0)

	errorhandler := func(err error) {
		log.Println("I am an error handler:", err)
	}

	preTaskHandler := func(signature *tasks.Signature) {
		log.Println("I am a start of task handler for:", signature.Name)
	}

	postTaskHandler := func(signature *tasks.Signature) {
		log.Println("I am an end of task handler for:", signature.Name)
	}

	worker.SetPostTaskHandler(postTaskHandler)
	worker.SetErrorHandler(errorhandler)
	worker.SetPreTaskHandler(preTaskHandler)

	return worker.Launch()
}
