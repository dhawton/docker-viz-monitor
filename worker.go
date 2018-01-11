package main

import (
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

var nodes []*Nodes
var services []*Services
var tasks []*Tasks

func worker(cli *client.Client) {
	startTime := time.Now().UnixNano()

	newnodes := make([]*Nodes, 0, 5)
	newservices := make([]*Services, 0, 5)
	newtasks := make([]*Tasks, 0, 5)

	mynodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}
	for _, node := range mynodes {
		n := &Nodes{
			ID:      node.ID,
			Name:    node.Description.Hostname,
			State:   string(node.Status.State),
			Role:    string(node.Spec.Role),
			Version: node.Description.Engine.EngineVersion,
			updated: startTime,
		}
		n.Tasks = make([]*Tasks, 0, 5)
		newnodes = append(newnodes, n)
	}
	nodes = newnodes

	myservices, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}
	for _, service := range myservices {
		s := &Services{
			ID:    service.ID,
			Name:  service.Spec.Name,
			Image: service.Spec.Labels["com.docker.stack.image"],
			updated: startTime,
		}

		newservices = append(newservices, s)
	}
	services = newservices

	mytasks, err := cli.TaskList(context.Background(), types.TaskListOptions{})
	if err != nil {
		panic(err)
	}
	for _, task := range mytasks {
		if task.DesiredState != swarm.TaskStateShutdown {
			t := &Tasks{
				ID:            task.ID,
				ServiceID:     task.ServiceID,
				Image:         task.Spec.ContainerSpec.Image,
				Hostname:      task.Spec.ContainerSpec.Hostname,
				NodeID:        task.NodeID,
				Status:        string(task.Status.State),
				DesiredStatus: string(task.DesiredState),
				updated: startTime,
			}
			newtasks = append(newtasks, t)
			findTaskOrAdd(t.NodeID, t)
		}
	}
	tasks = newtasks
}

func initWorker() {
	os.Setenv("DOCKER_API_VERSION", "1.35")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	for {
		<-time.After(2 * time.Second)
		worker(cli)
	}
}
