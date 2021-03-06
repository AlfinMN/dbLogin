package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func ReadEnv(key, defVal string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Keyname : %v, not found, default key value : %v, has been loaded", key, defVal)
		//log.Println("Default Config has been loaded")
		return defVal
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func Connection() (*sql.DB, error) {

	dbUser := ReadEnv("dbUser", "root")
	dbPass := ReadEnv("dbPass", "Apinchocs98")
	dbHost := ReadEnv("dbHost", "localhost")
	dbPort := ReadEnv("dbPort", "3306")
	dbName := ReadEnv("dbName", "login")

	loadData := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open("mysql", loadData)

	err := db.Ping()
	if err != nil {
		log.Print(err)
	}

	return db, nil
}
