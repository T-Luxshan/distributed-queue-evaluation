package server

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"log"
)

func StartServer() (*machinery.Server, error) {
	cnf := &config.Config{
		Broker:          "amqp://guest:guest@localhost:5672/",
		DefaultQueue:    "process_tasks",
		ResultBackend:   "amqp://guest:guest@localhost:5672/",
		ResultsExpireIn: 3600,
		AMQP: &config.AMQPConfig{
			Exchange:      "machinery_exchange",
			ExchangeType:  "direct",
			BindingKey:    "ride_task",
			PrefetchCount: 3,
		},
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}

	// Register the task.
	tasks := map[string]interface{}{
		"process_tasks": processTask,
	}
	return server, server.RegisterTasks(tasks)
}

func processTask(userID, rideID string) error {

	log.Printf("Processing ride: userID=%s, rideID=%s,\n", userID, rideID)
	if userID == "" || rideID == "" {
		return fmt.Errorf("invalid ride data: userID=%d, rideID=%d, fare=%.2f", userID, rideID)
	}
	fmt.Printf("Ride processed successfully: userID=%s, rideID=%s", userID, rideID)
	return nil
}
