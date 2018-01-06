package main

import (
	"os"
	"time"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

var nodes map[string]Nodes
var services map[string]Services
var tasks map[string]Tasks

func worker(cli *client.Client) {
	mynodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}
	for _, node := range mynodes {
		mynode := Nodes{
			ID: node.ID,
			Name: node.Spec.Labels["short"],
			State: string(node.Status.State),
			Role: string(node.Spec.Role),
			Version: node.Description.Engine.EngineVersion,
		}
		nodes[mynode.ID] = mynode
	}

	myservices, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}
	for _, service := range myservices {
		myservice := Services{
			ID: service.ID,
			Name: service.Spec.Name,
			Image: service.Spec.Labels["com.docker.stack.image"],
		}

		services[myservice.ID] = myservice
	}

	mytasks, err := cli.TaskList(context.Background(), types.TaskListOptions{})
	if err != nil {
		panic(err)
	}
	for _, task := range mytasks {
		t := Tasks{
			ID: task.ID,
			ServiceID: task.ServiceID,
			NodeID: task.NodeID,
			Status: string(task.Status.State),
			DesiresStatus: string(task.DesiredState),
			Node: nodes[task.NodeID],
		}

		tasks[t.ID] = t
	}
}

func initWorker() {
	os.Setenv("DOCKER_API_VERSION", "1.35");
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	nodes = make(map[string]Nodes)
	services = make(map[string]Services)
	tasks = make(map[string]Tasks)

	for {
		<- time.After(2 * time.Second)
		worker(cli)
	}
}