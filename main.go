package main

import (
	"PromotionsAPI/initializers"
	"PromotionsAPI/routers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	routers.ConfigureRoutes()
}
