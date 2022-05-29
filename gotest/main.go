package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("ONg", 10, nil)
	c.On("GetCustomer", 2).Return("", 10, errors.New("not found"))

	fmt.Println(c.GetCustomer(1))
	fmt.Println(c.GetCustomer(2))
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2)
}
