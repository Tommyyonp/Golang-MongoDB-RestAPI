package main

import (
	"Golang-MongoDB-RestAPI/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	// define router
	router := httprouter.New()
	userController := controllers.NewUserController(getSession())

	router.GET("/user/:id", userController.GetUser)
	router.POST("/user", userController.CreateUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	return session
}
