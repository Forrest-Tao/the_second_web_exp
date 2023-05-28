package main

import (
	"experimen_2/conf"
	"experimen_2/router"
)

func main() {
	conf.Init()

	engine := router.NewRouter()
	engine.Run(conf.HttpPort)

}
