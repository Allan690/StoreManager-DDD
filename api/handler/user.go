package handler

import (
	"StoreManager-DDD/api/presenter"
	"StoreManager-DDD/entity"
	"StoreManager-DDD/usecase/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func listUsers(service user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userList, err := service.ListUsers()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "users": &userList})

	}
}

func createUser(service user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Email     string `json:"email"`
			Password  string `json:"password"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}
		err := c.Bind(&input)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}
		id, err := service.CreateUser(input.Email, input.Password, input.FirstName, input.LastName)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}
		toJ := &presenter.User{
			ID:        entity.ID{UUID: id},
			Email:     input.Email,
			FirstName: input.FirstName,
			LastName:  input.LastName,
		}
		c.JSON(http.StatusCreated, gin.H{"status": "success", "user": toJ})
	}
}

func MakeUserHandlers(r *gin.Engine, service user.UseCase) *gin.Engine {
	api := r.Group("/api")

	{
		api.GET("/user",  listUsers(service))
		api.POST("/user", createUser(service))
	}
	return r
}
