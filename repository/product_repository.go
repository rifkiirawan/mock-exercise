package repository

import "auth/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
	FindAll() []*entity.Product
}
