package repository

import (
	"github.com/ko-he/test-goloan-sample/domain/model"
)

// ClientRepository interface
type ClientRepository interface {
	Featch() ([]*model.Client, error)
}
