package main

import (
	"crud/controllers"
	"crud/initialize"
)

func main() {
	init := &initialize.AppSettings{}
	cs, p := init.Init("dev")
	controllers.Start()
	_, _ = cs, p

}
