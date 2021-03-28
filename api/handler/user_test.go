package handler_test

import (
	"StoreManager-DDD/api/handler"
	"StoreManager-DDD/entity"
	mock "StoreManager-DDD/usecase/user/mock"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_createUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := gin.Default()
	r = handler.MakeUserHandlers(r, m)
	m.EXPECT().
		CreateUser(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(entity.NewID(), nil)
	handler.CreateUser(m)
	w := httptest.NewRecorder()
	payload := fmt.Sprintf(`{
		"email": "ozzy@hell.com",
		"password": "asasa",
		"first_name":"Ozzy",
		"last_name":"Osbourne"
		}`)
	req, err := http.NewRequest("POST", "/api/user", strings.NewReader(payload))
	r.ServeHTTP(w, req)
	sec := map[string]interface{}{}
	_ = json.Unmarshal([]byte(w.Body.String()), &sec)
	assert.Nil(t, err)
	fmt.Println(sec)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, sec["status"], "success")
	uuidString := (sec["user"]).(map[string]interface{})["id"]
	assert.IsType(t, uuidString, entity.ID{UUID: entity.NewID()}.String(), nil)
}

func TestListUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := gomock.NewController(t)
	defer controller.Finish()
	r := gin.Default()
	m := mock.NewMockUseCase(controller)
	handler.MakeUserHandlers(r, m)
	u := &entity.User{
		ID: entity.ID{UUID: entity.NewID()},
	}
	m.EXPECT().
		ListUsers().
		Return([]*entity.User{u}, nil)
	handler.ListUsers(m)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/user", nil)
	r.ServeHTTP(w, req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := gomock.NewController(t)
	defer controller.Finish()
	r := gin.Default()
	m := mock.NewMockUseCase(controller)
	handler.MakeUserHandlers(r, m)
	u := &entity.User{
		ID: entity.ID{UUID: entity.NewID()},
		Email: "test@test.com",
	}
	m.EXPECT().GetUser(u.ID.UUID).Return(u, nil)
	m.EXPECT().UpdateUser(u).Return(nil)
	handler.UpdateUser(m)
	w := httptest.NewRecorder()
	payload, _ := json.Marshal(map[string]interface{} {
		"email": "testing@testing.com",
	})
	req, err := http.NewRequest(
		"PATCH",
		fmt.Sprintf("/api/user/%s",
			u.ID.UUID.String()),
			bytes.NewBuffer(payload),
		)
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Header())
}