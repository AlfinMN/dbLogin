package usecase

import "testdong/login/master/model"

type LoginUsecase interface {
	GetDataLogin() ([]*model.Login, error)
}
