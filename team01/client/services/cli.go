package services

import (
	"errors"
	"strings"
	"warehouse/client/internal/models"
)

func GetCommand(input string) (models.ClientRequest, error) {
	input = strings.TrimSpace(input)

	args := strings.Split(input, " ")
	if len(args) < 2 {
		return models.ClientRequest{}, errors.New("invalid command")
	}

	command := strings.ToUpper(args[0])
	var value string
	if len(args) == 3 {
		value = args[3]
	}

	if command == "SET" || command == "GET" || command == "DELETE" {
		return models.ClientRequest{
			Command: args[0],
			Key:     args[1],
			Value:   value,
		}, nil
	}

	return models.ClientRequest{}, errors.New("invalid command - expected [GET, SET, DELETE]")
}
