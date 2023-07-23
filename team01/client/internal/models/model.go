package models

type Heartbeat struct {
	Command   string `json:"command"`
	ClientURL string `json:"client_address"`
}

type HeartbeatResponse struct {
	Command           string `json:"command"`
	Status            string `json:"status"`
	ReplicationFactor int    `json:"replication_factor"`
	Nodes             []Node `json:"nodes"`
}

type Node struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type ClientRequest struct {
	Command string `json:"command"`
	Key     string `json:"key"`
	Value   string `json:"value,omitempty"`
}
