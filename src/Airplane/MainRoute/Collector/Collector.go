package Collector

import (
	"Cargo/Blog/route"
	"Airplane/MainRoute"
)

func Collector()  {
	route.BlogRoute()

	//custom Route
	MainRoute.Airplane.Run()
}