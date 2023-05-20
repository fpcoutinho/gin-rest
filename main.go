package main

import (
	"github.com/fpcoutinho/gin-rest/configs"
	"github.com/fpcoutinho/gin-rest/routes"
)

func main() {
	configs.ConnectDB()
	routes.HandleRequests()
}
