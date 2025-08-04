package main

import (
	"usdrive/config"
	"usdrive/routes"
)

func main() {
	router := routes.Master()
	router.Run(":"+config.ENV.Port)
}

