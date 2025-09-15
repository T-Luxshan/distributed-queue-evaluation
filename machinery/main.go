package main

import (
	"fmt"
	myTask "github.com/T-Luxshan/distributed-queue-evaluation/machinery/tasks"
	"github.com/T-Luxshan/distributed-queue-evaluation/machinery/worker"
	"github.com/urfave/cli"
	"os"
)

var (
	app *cli.App
)

func init() {
	// Initialize a CLI app
	app = cli.NewApp()
	app.Name = "machinery"
	app.Usage = "machinery worker and send ride details from CSV"
	app.Version = "0.0.0"
}

func main() {
	// Set the CLI app commands
	app.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "launch machinery worker",
			Action: func(c *cli.Context) error {
				if err := worker.Worker(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
		{
			Name:  "enqueue",
			Usage: "enqueue ride details from CSV",
			Action: func(c *cli.Context) error {
				if err, _ := myTask.Enqueue(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
		{
			Name:  "peek",
			Usage: "peek ride details from queue",
			Action: func(c *cli.Context) error {
				if err, _ := myTask.PeekMessage(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
		{
			Name:  "dequeue",
			Usage: "dequeue ride details from queue",
			Action: func(c *cli.Context) error {
				if err, _ := myTask.DequeueMessage(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
		{
			Name:  "eval",
			Usage: "Evaluate the queue operations",
			Action: func(c *cli.Context) error {
				evaluate()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}

func evaluate() {

	_, t1 := myTask.Enqueue()
	_, t2 := myTask.PeekMessage()
	_, t3 := myTask.DequeueMessage()

	fmt.Printf("\n\n elapsedEnqueue: %s\n elapsedPeak: %s\n elapsedDequeue: %s\n", t1, t2, t3)

}
