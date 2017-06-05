package Collector

import (
	"Cargo/Blog/route"
	"Airplane/MainRoute"
)

func Collector()  {
	route.BlogRoute()


	MainRoute.Airplane.Run()
}