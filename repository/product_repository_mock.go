package repository

import (
	"auth/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *entity.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(entity.Product)

	return &product
}

func (repository *ProductRepositoryMock) FindAll() []*entity.Product {
	arguments := repository.Mock.Called()

	if arguments == nil {
		return nil
	}

	arguments0 := arguments.Get(0)

	if arguments0 == nil {
		return nil
	}

	products := arguments0.([]*entity.Product)

	// var products []*entity.Product
	// for _, v := range arguments {
	// 	products = append(products, &v)
	// }

	return products
}
