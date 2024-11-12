package routes

import (
	"chatbotIA/controllers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/",
	controllers.HomeController)
}