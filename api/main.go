package main

import (
	"StoreManager-DDD/api/handler"
	"StoreManager-DDD/config"
	"StoreManager-DDD/infrastructure"
	"StoreManager-DDD/usecase/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/mgosession"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	session, err := mgo.Dial(config.DB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()
	mPool := mgosession.NewPool(nil, session, 5)
	defer mPool.Close()

	userRepo := infrastructure.NewMongoRepository(mPool)
	userService := user.NewService(userRepo)
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	handler.MakeUserHandlers(router, userService)
	err = router.Run(":8000")
	if err != nil {
		fmt.Print(err)
		panic("An error occurred when running this application")
	}
}
