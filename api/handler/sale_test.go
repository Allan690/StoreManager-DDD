package handler_test

import (
	"StoreManager-DDD/api/handler"
	"StoreManager-DDD/entity"
	mock "StoreManager-DDD/usecase/sale/mock"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var mockProduct = uuid.Must(uuid.NewV4())
var mockUser = uuid.Must(uuid.NewV4())
var mockSale = entity.Sale{
	ID:        entity.ID{UUID: entity.NewID()},
	Product:   entity.ID{UUID: mockProduct},
	User:      entity.ID{UUID: mockUser},
	Total:     100,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}
var mockSalesList []*entity.Sale

func TestCreateSales(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUsecase(controller)
	r := gin.Default()
	r = handler.MakeSalesHandlers(r, m)
	m.EXPECT().
		CreateSale(gomock.Any()).
		Return(&mockSale, nil)
	handler.CreateSales(m)
	w := httptest.NewRecorder()
	payload := fmt.Sprintf( `{
		"product": %s,
		"user": %s,
		"total": %d
		}`,uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), 500)
	req, err := http.NewRequest("POST", "/api/sale", strings.NewReader(payload))
	r.ServeHTTP(w, req)
	sec := map[string]interface{}{}
	_ = json.Unmarshal([]byte(w.Body.String()), &sec)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, sec["status"], "success")
}

func TestGetSalesByUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUsecase(controller)
	r := gin.Default()
	r = handler.MakeSalesHandlers(r, m)
	mockSalesList = append(mockSalesList, &mockSale, &mockSale)
	m.EXPECT().
		GetSalesByUserId(mockSale.User.UUID).
		Return(mockSalesList, nil)
	handler.GetSalesByUserId(m)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/sale/user/%s", mockSale.User.UUID), nil)
	r.ServeHTTP(w, req)
	sec := map[string]interface{}{}
	_ = json.Unmarshal([]byte(w.Body.String()), &sec)
	fmt.Println(sec)
	assert.Nil(t, err)
	assert.Len(t, sec, 2)
	assert.Equal(t, http.StatusOK, w.Code)
}

//func TestListSales(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//	m := mock.NewMockUsecase(controller)
//	r := gin.Default()
//	r = handler.MakeSalesHandlers(r, m)
//	m.EXPECT().GetAllSales().Return(&mockSale, nil)
//}