package route

import "Airplane/MainRoute"

func BlogRoute()  {
	MainRoute.Airplane.Get("/", func() (string) {
		return "Hello"
	})
}
