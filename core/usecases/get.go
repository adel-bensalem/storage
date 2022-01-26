package usecases

import "github.com/adel-bensalem/storage/core/adapters"

type GetUseCase interface {
	get(key string) (string, error)
}

type GetInteractor struct {
	repository *adapters.Repository
}

func (interactor GetInteractor) get(key string) (string, error) {
	return interactor.get(key)
}
