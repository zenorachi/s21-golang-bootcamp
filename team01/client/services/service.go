package services

import (
	"encoding/json"
	"net"
	"warehouse/client/internal/models"
)

func GetHeartBeat(addr string) (models.HeartbeatResponse, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return models.HeartbeatResponse{}, err
	}
	defer conn.Close()

	urlClient := conn.LocalAddr().String()
	heartbeat := models.Heartbeat{
		Command:   "heartbeat",
		ClientURL: urlClient,
	}

	encoder := json.NewEncoder(conn)
	if err = encoder.Encode(heartbeat); err != nil {
		return models.HeartbeatResponse{}, err
	}

	var heartbeatResp models.HeartbeatResponse
	decoder := json.NewDecoder(conn)
	if err = decoder.Decode(&heartbeatResp); err != nil {
		return models.HeartbeatResponse{}, err
	}

	return heartbeatResp, nil
}
