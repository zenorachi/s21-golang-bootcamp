package config

import (
	"encoding/json"
	"os"
)

type candy struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type Candies struct {
	Models []candy `json:"candies"`
}

func (c *Candies) FillModels() error {
	data, err := os.ReadFile("./candy-server/config/config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Candies) GetMapModels() map[string]int64 {
	candyMap := make(map[string]int64)
	for _, model := range c.Models {
		candyMap[model.Name] = model.Price
	}
	return candyMap
}
