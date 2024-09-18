package main

import (
	"chi-onion-arch/app"
	"chi-onion-arch/config"
)

func main() {
	config.Init()
	app.Run()
}
