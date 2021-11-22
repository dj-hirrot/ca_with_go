package usecase

import "github.com/dj-hirrot/ca_with_go/src/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	return
}

func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) CreateUser(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) UpdateUser(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Update(u)
	return
}

func (interactor *UserInteractor) DeleteUser(u domain.User) (err error) {
	err = interactor.UserRepository.DeleteById(u)
	return
}
