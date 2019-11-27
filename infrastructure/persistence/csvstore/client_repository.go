package csvstore

import (
	"encoding/csv"
	"os"

	"github.com/ko-he/test-goloan-sample/domain/model"
	"github.com/ko-he/test-goloan-sample/domain/repository"
)

type clientRepository struct{}

// NewClientRepository func
func NewClientRepository() repository.ClientRepository {
	return &clientRepository{}
}

func (c *clientRepository) Featch() ([]*model.Client, error) {
	clients := []*model.Client{}
	file, err := os.Open("./companies.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	i := 0
	var line []string

	for {
		line, err = reader.Read()
		if err != nil {
			break
		}

		i++

		client := &model.Client{}
		client.ID = i
		client.Name = line[0]
		client.Address = line[1]

		clients = append(clients, client)
	}
	return clients, err
}
