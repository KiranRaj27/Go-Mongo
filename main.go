package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Kiranraj27/mongo-go/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	router := httprouter.New()
	uc := controllers.NewUserController(getSession())
	router.GET("/user/:id", uc.Getuser)
	router.POST("/user", uc.CreateUser)

	http.ListenAndServe("localhost:8000", router)
}

func getSession() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:kichu123@myatlasclusteredu.5ekwdbt.mongodb.net/golang"))
	if err != nil {
		fmt.Println(err)
	}
	return client
}
