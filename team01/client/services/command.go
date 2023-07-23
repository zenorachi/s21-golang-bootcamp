package services

import (
	"encoding/json"
	"net"
	"warehouse/client/internal/models"
)

func SendRequestCommand(addr string, command models.ClientRequest) (map[string]interface{}, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	err = encoder.Encode(command)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(conn)
	var response map[string]interface{}
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
