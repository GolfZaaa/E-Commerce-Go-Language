package middlewaresUsecases

import (
	"github.com/GolfZaaa/Golang-ProjectE-commerce/modules/middlewares/middlewaresRepositories"
)

type IMiddlewaresUsecase interface {
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewaresRepository
}

func MiddlewareUsecase(middlewaresRepository middlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}
