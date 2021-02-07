package main

import (
	"backend-shortcourse/go-webservice/config"
	"backend-shortcourse/go-webservice/database"
	"backend-shortcourse/go-webservice/server"
)

func main() {
	config.Init()
	database.Init()
	server.Init()
}
