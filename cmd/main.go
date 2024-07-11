package main

import "api_gateway/api"

func main() {

	r := api.NewRouter()
	r.Run()

}
