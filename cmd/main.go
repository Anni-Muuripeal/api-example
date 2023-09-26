package main

import (
	"api-example/pkg/api"
)

func main() {
	server := api.ApiServer{}
	server.Run()
}
