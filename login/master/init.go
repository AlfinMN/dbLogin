package master

import (
	"testdong/config"
	"testdong/login/master/controller"
	"testdong/login/master/repositories"
	"testdong/login/master/usecase"

	"github.com/gorilla/mux"
)

func Init(izipizy *mux.Router) {

	db, _ := config.Connection()
	repoLogin := repositories.InitLoginRepoImpl(db)
	usecaseLogin := usecase.InitUsecaseImpl(repoLogin)
	controller.GetDataController(izipizy, usecaseLogin)
}
