package handler

import (
	"StoreManager-DDD/api/presenter"
	"StoreManager-DDD/entity"
	"StoreManager-DDD/usecase/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"strings"
	"time"
)

func ListUsers(service user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userList, err := service.ListUsers()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "users": &userList})

	}
}

func CreateUser(service user.UseCase) gin.HandlerFunc {
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

func UpdateUser(service user.UseCase) gin.HandlerFunc {
	type UserUpdate struct {
		Email string `json:"email,omit_empty"`
		FirstName string `json:"first_name,omit_empty"`
		LastName string `json:"last_name,omit_empty"`
	}
	return func(c *gin.Context) {
		var id = c.Param("id")
		existingUser, err := service.GetUser(uuid.Must(uuid.FromString(id)))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
			return
		}
		var userUpdate UserUpdate
		err = c.Bind(&userUpdate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}
		fmt.Println(userUpdate.Email)
		if len(strings.TrimSpace(userUpdate.FirstName)) > 0 {
			existingUser.FirstName = userUpdate.FirstName
		}
		if len(strings.TrimSpace(userUpdate.LastName)) > 0 {
			existingUser.LastName = userUpdate.LastName
		}
		if len(strings.TrimSpace(userUpdate.Email)) > 0 {
			existingUser.Email = userUpdate.Email
		}
		existingUser.UpdatedAt = time.Now()
		err = service.UpdateUser(existingUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "user": &existingUser})
	}
}

func MakeUserHandlers(r *gin.Engine, service user.UseCase) *gin.Engine {
	api := r.Group("/api")

	{
		api.GET("/user",  ListUsers(service))
		api.POST("/user", CreateUser(service))
		api.PATCH("/user/:id", UpdateUser(service))
	}
	return r
}
