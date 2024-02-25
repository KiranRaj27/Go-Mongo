package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kiranraj27/mongo-go/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	fmt.Println("hello")
	router := httprouter.New()
	uc := controllers.NewUserController(getSession())
	router.GET("/user/:id", uc.Getuser)
	router.POST("/user", uc.CreateUser)

	http.ListenAndServe("localhost:8000", router)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb+srv://admin:kichu123@myatlasclusteredu.5ekwdbt.mongodb.net/golang")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Something went wrong")
	}
	return s
}
