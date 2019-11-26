package service

import (
	"strings"

	"github.com/ko-he/test-goloan-sample/domain/model"
	"github.com/ko-he/test-goloan-sample/domain/repository"
)

// ClientService interface
type ClientService interface {
	SearchByPrefacture(p string) ([]*model.Client, error)
}

type clientService struct {
	repository.ClientRepository
}

// NewClientService func
func NewClientService(r repository.ClientRepository) ClientService {
	return &clientService{r}
}

func (c *clientService) SearchByPrefacture(p string) ([]*model.Client, error) {
	clients, _ := c.ClientRepository.Featch()

	var selectedClients []*model.Client

	for _, c := range clients {
		if strings.HasPrefix(c.Address, p) {
			selectedClients = append(selectedClients, c)
		}
	}

	return selectedClients, nil
}
