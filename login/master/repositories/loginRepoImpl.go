package repositories

import (
	"database/sql"
	"log"
	"testdong/login/master/model"
)

type LoginRepoImpl struct {
	db *sql.DB
}

func (database LoginRepoImpl) GetDataUser() ([]*model.Login, error) {
	var data []*model.Login
	loginUser, err := database.db.Query("select * from user ")
	if err != nil {
		log.Fatal(err)
	}
	for loginUser.Next() {
		userLogin := model.Login{}
		err := loginUser.Scan(&userLogin.ID, &userLogin.Username, &userLogin.Paswword)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, &userLogin)
	}
	return data, nil
}

func InitLoginRepoImpl(db *sql.DB) LoginRepo {
	return &LoginRepoImpl{db}
}
