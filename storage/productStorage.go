package storage

import "pos.com/app/domain"

type ProductStorage struct{}

func (s ProductStorage) Get(id int) *domain.Product {
	return &domain.Product{
		Name: "",
	}
}
