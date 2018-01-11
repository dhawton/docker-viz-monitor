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
	for i := range nodes {
		if nodes[i].ID == nodeID {
			for _, v := range nodes[i].Tasks {
				if v == task.ID {
					found = true
				}
			}
			if found == false {
				(nodes[i]).Tasks = append((nodes[i]).Tasks, task.ID)
			}
		}
	}
}