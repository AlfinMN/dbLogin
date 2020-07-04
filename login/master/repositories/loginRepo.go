package repositories

import "testdong/login/master/model"

type LoginRepo interface {
	GetDataUser() ([]*model.Login, error)
}
