package main

import (
	"crud/controllers"
)

func main() {
	controller := controllers.Controllers{}
	controller.Start()
}
