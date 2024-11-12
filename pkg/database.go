package pkg

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

// ConnectDatabase establece la conexion con MongoDb

func ConnectDatabase() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL no esta configurada en el archivo .env")
	}
	
	clientOptions := options.Client().ApplyURI(dbURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	// Verificar la conexion
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("No se pudo establecer la conexion a la base de datos:", err)
	}

	log.Println("Conexion a la base de datos exitosa")
	DB = client
}

