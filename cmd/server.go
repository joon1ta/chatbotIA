package main

import (
	"chatbotIA/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	// OBtener el puerto del servidor desde las variables de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //Valor predeterminado si no se encuentra el puerto en .env
	}

	//Registrar las rutas
	routes.RegisterRoutes()

	// Iniciar el servidor
	log.Printf("Server staring on port %s/n", port)
	if err := http.ListenAndServe(":" + port, nil); err !=
	nil {
		log.Fatal(err)
	}
}