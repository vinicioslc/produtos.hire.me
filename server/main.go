package main

import (
	"bemobi-hire/server/api"
	"bemobi-hire/server/conf"
	"fmt"
)

func main() {
	apiconfig := conf.LoadAppConfig()
	fmt.Println("RUNING API ON PORT:", apiconfig.AppPort)
	fmt.Println(apiconfig.AppName, "SERVER - STARTED")
	// Inicia a API
	api.RunAPI(apiconfig)

}
