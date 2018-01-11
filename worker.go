package main

import (
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

var nodes map[string]*Nodes
var services map[string]Services
var tasks map[string]Tasks

func worker(cli *client.Client) {
	startTime := time.Now().UnixNano()

	mynodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}
	for _, node := range mynodes {
		if nodes[node.ID] == nil {
			nodes[node.ID] = &Nodes{
				ID:      node.ID,
				Name:    node.Spec.Labels["short"],
				State:   string(node.Status.State),
				Role:    string(node.Spec.Role),
				Version: node.Description.Engine.EngineVersion,
				updated: startTime,
			}
			nodes[node.ID].Tasks = make([]string, 0, 5)
		} else {
			nodes[node.ID].Name = node.Spec.Labels["short"]
			nodes[node.ID].State = string(node.Status.State)
			nodes[node.ID].Role = string(node.Spec.Role)
			nodes[node.ID].Version = node.Description.Engine.EngineVersion
			nodes[node.ID].updated = startTime
		}
	}
	removeExpiredNodes(startTime)

	myservices, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}
	for _, service := range myservices {
		myservice := Services{
			ID:    service.ID,
			Name:  service.Spec.Name,
			Image: service.Spec.Labels["com.docker.stack.image"],
		}

		services[myservice.ID] = myservice
	}

	mytasks, err := cli.TaskList(context.Background(), types.TaskListOptions{})
	if err != nil {
		panic(err)
	}
	for _, task := range mytasks {
		if task.DesiredState != swarm.TaskStateShutdown {
			t := Tasks{
				ID:            task.ID,
				ServiceID:     task.ServiceID,
				Image:         task.Spec.ContainerSpec.Image,
				Hostname:      task.Spec.ContainerSpec.Hostname,
				NodeID:        task.NodeID,
				Status:        string(task.Status.State),
				DesiredStatus: string(task.DesiredState),
			}
			tasks[t.ID] = t
			findTaskOrAdd(t.NodeID, t)
		}
	}
}

func initWorker() {
	os.Setenv("DOCKER_API_VERSION", "1.35")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	nodes = make(map[string]*Nodes)
	services = make(map[string]Services)
	tasks = make(map[string]Tasks)

	for {
		<-time.After(2 * time.Second)
		worker(cli)
	}
}
