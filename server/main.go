package main

import (
	"bemobi-hire/server/api"
	"bemobi-hire/server/config"
	"fmt"
)

func main() {
	apiconfig := config.LoadAppConfig()
	fmt.Println("RUNING API ON PORT:", apiconfig.AppPort)
	fmt.Println(apiconfig.AppName, "SERVER - STARTED")
	// Inicia a API
	api.RunAPI(apiconfig)

}
