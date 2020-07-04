package usecase

import (
	"testdong/login/master/model"
	"testdong/login/master/repositories"
)

type LoginUsecaseImpl struct {
	LoginRepoImpl repositories.LoginRepo
}

func (user LoginUsecaseImpl) GetDataLogin() ([]*model.Login, error) {
	data, err := user.LoginRepoImpl.GetDataUser()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InitUsecaseImpl(LoginRepoImpl repositories.LoginRepo) LoginUsecase {
	return &LoginUsecaseImpl{LoginRepoImpl}
}
