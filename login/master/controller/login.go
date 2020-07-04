package controller

import (
	"encoding/json"
	"net/http"
	"testdong/login/master/model"
	"testdong/login/master/usecase"

	"github.com/gorilla/mux"
)

type LoginHandler struct {
	loginUsecase usecase.LoginUsecase
}

func GetDataController(ez *mux.Router, dataUsecase usecase.LoginUsecase) {
	handlerLogin := LoginHandler{dataUsecase}
	ez.HandleFunc("/login", handlerLogin.LoginUser).Methods("GET")
}

func (lh LoginHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	dataUser, _ := lh.loginUsecase.GetDataLogin()

	var pesan model.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "data login succes displayed :) "
	pesan.Data = dataUser

	dataLogin, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataLogin)

}
