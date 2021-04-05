package handler

import (
	"StoreManager-DDD/entity"
	"StoreManager-DDD/usecase/sale"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"time"
)

var CreateSaleStruct struct {
	Product uuid.UUID `json:"product"`
	User uuid.UUID `json:"user"`
	Total int `json:"total"`
}


// ListSales handles listing of sales
func ListSales(service sale.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userList, err := service.GetAllSales()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "sales": &userList})
	}
}

// CreateSales handles creation of a sale
// returns a handler function that processes the http request
func CreateSales(service sale.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.Bind(&CreateSaleStruct)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}
		sale := entity.Sale{}
		sale.ID = entity.ID{UUID: entity.NewID()}
		sale.Product = entity.ID{UUID: CreateSaleStruct.Product}
		sale.User = entity.ID{UUID: CreateSaleStruct.User}
		sale.CreatedAt = time.Now()
		sale.UpdatedAt = time.Now()
		newSale, err := service.CreateSale(&sale)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"status": "success", "sale": newSale})
	}
}

//GetSalesByUserId gets sales by a certain user using their id
func GetSalesByUserId(service sale.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "invalid id"})
			return
		}
		sales, err := service.GetSalesByUserId(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "sales": sales})
	}
}

// DeleteSale handles creation of a sale and returns a handler function that
//handles the incoming http request.
// the sale service is injected as a dependency
func DeleteSale(service sale.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "invalid id"})
			return
		}
		_, err = service.DeleteSale(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{"status": "success",
			"message": fmt.Sprintf( "sale deleted!")})
	}
}


func MakeSalesHandlers(r *gin.Engine, service sale.Usecase) *gin.Engine {
	api := r.Group("/api")

	{
		api.GET("/sale",  ListSales(service))
		api.POST("/sale", CreateSales(service))
		api.GET("/sale/user/:id", GetSalesByUserId(service))
		api.DELETE("/sale/:id", DeleteSale(service))
	}
	return r
}
