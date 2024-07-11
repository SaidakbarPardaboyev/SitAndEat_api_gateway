package main

import (
	"api_gateway/api"
	"api_gateway/config"
)

func main() {

	r := api.NewRouter()
	r.Run(config.Load().API_GATEWAY)

}
