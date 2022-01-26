package core

import "github.com/adel-bensalem/storage/core/usecases"

type Core interface {
	usecases.GetUseCase
}

type Interactor struct {
	usecases.GetInteractor
}
