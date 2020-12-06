package service

import "wire_sample/repository"

type (
	Service struct {
		Repository repository.Repository
	}
)

func NewService(repo repository.Repository) Service {
	return Service{repo}
}
