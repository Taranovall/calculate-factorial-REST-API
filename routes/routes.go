package routes

import (
	"calculateFactorial/handlers"
	"calculateFactorial/middlewars"
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{"CalculateFactorial", "POST", "/calculate", middlewars.ValidateFactorialRequest(handlers.CalculateFactorial)},
	}
	return routes
}
