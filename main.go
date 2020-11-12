package main

import (
	"context"
	"log"
	
	userHandler "mongodb-api/user/handler"
	userRepo "mongodb-api/user/repo"
	userUsecase "mongodb-api/user/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
	var port = "8098"
	var ctx = context.Background()
	db, err := connectDB(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	userRepo := userRepo.CreateUserRepoImpl(db, ctx)
	userUsecase := userUsecase.CreateUserUsecase(userRepo)

	userHandler.CreateUserHandler(router, userUsecase)
	
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err.Error())
	}
}



func connectDB(ctx context.Context) (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client.Database("belajar_golang_api"), nil
}