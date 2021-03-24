package handler

import (
	"StoreManager-DDD/entity"
	mock "StoreManager-DDD/usecase/user/mock"
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
	r = MakeUserHandlers(r, m)
	m.EXPECT().
		CreateUser(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(entity.NewID(), nil)
	createUser(m)
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

