package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

type Nodes struct {
	ID    string
	Name  string
	State string
	Role  string
	Version string
}

type Services struct {
	ID    string
	Name  string
	Image string
}

type Tasks struct {
	ID            string
	ServiceID     string
	NodeID        string
	Status        string
	DesiresStatus string
}

func worker(cli *client.Client) {
	nodes := make(map[string]Nodes)
	services := make(map[string]Services)
	tasks := make(map[string]Tasks)
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
		}

		tasks[t.ID] = t
	}

	nj, _ := json.Marshal(nodes)
	sj, _ := json.Marshal(services)
	tj, _ := json.Marshal(tasks)

	err = ioutil.WriteFile(os.Getenv("JSON_NODES"), nj, 0644)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(os.Getenv("JSON_SERVICES"), sj, 0644)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(os.Getenv("JSON_TASKS"), tj, 0644)
}

func main() {
	os.Setenv("DOCKER_API_VERSION", "1.35");
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	for {
		<-time.After(2 * time.Second)
		go worker(cli)
	}
}
