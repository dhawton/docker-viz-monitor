package main

type Nodes struct {
	ID      string
	Name    string
	State   string
	Role    string
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
	Node          Nodes
}
