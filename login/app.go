package main

import (
	"testdong/config"
	"testdong/login/master"
)

func main() {

	routing := config.CreateRouter()
	master.Init(routing)
	config.RunServer(routing)
}
