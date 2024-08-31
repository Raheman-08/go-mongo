package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/Raheman-08/go-mongo/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017/")
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		panic(err)
	}
	fmt.Println("Connected to MongoDB")
	return s
}
