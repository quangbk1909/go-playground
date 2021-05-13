package main

import (
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	machineryConfig "github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"time"
)

func NewMachineryServer(taskServerBroker, taskServerQueueDefault, taskServerResultBackend string) *machinery.Server {
	var cnf = &machineryConfig.Config{
		Broker:        taskServerBroker,
		DefaultQueue:  taskServerQueueDefault,
		ResultBackend: taskServerResultBackend,
	}
	server, err := machinery.NewServer(cnf)
	if err != nil {
		panic(err)
	}
	return server
}

func main() {
	redisLocal := "redis://127.0.0.1:6379"
	taskQueue := "abc"
	serverSender := NewMachineryServer(redisLocal, taskQueue, redisLocal)

	data :=
	jsonData, _ := json.Marshal(data)

	signature := tasks.Signature{
		Name: "hehe",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: string(jsonData),
			},
		},
	}


	_, err := serverSender.SendTask(&signature)
	if err != nil {
		panic(err)
	}

	fmt.Println("send task success")


	time.Sleep(20 * time.Second)
}

func process(data string) error {
	time.Sleep(10*time.Second)
	fmt.Println(data)
	return nil
}

func runWorkerInstance(server *machinery.Server, tagName string) error {
	worker := server.NewWorker(tagName, 2)
	err := worker.Launch()
	if err != nil {
		return err
	}
	return nil
}




