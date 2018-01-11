package main

type Nodes struct {
	ID      string
	Name    string
	State   string
	Role    string
	Version string
	Tasks   []string
	updated int64
}

type Services struct {
	ID      string
	Name    string
	Image   string
	updated int64
}

type Tasks struct {
	ID            string
	Image         string
	Hostname      string
	ServiceID     string
	NodeID        string
	Status        string
	DesiredStatus string
	updated       int64
}

func findTaskOrAdd(nodeID string, task *Tasks) {
	found := false
	for _, v := range nodes[nodeID].Tasks {
		if v == task.ID {
			found = true
		}
	}
	if found == false {
		(nodes[nodeID]).Tasks = append((nodes[nodeID]).Tasks, task.ID)
	}
}

func removeExpiredNodes(stamp int64) {
	for k := range nodes {
		if nodes[k].updated != stamp {
			nodes[k] = nil
		}
	}
}

func removeExpiredServices(stamp int64) {
	for k := range services {
		if services[k].updated != stamp {
			services[k] = nil
		}
	}
}

func removeExpiredTasks(stamp int64) {
	for k := range tasks {
		if tasks[k].updated != stamp {
			tasks[k] = nil
		}
	}
}