package main

import (
	"context"
	"log"

	"github.com/wafi/hello-workflow/helloworkflow"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to make client", err)
	}

	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: "hello-world",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, helloworkflow.Workflow, "victor")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	var result string
	// store the result of the run
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}
	log.Println("workflow result:", result)
}
