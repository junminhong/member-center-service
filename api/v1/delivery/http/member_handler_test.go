package http_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	memberDeliver "github.com/junminhong/member-center-service/api/v1/delivery/http"
	"github.com/junminhong/member-center-service/domain/mocks"
	"github.com/junminhong/member-center-service/pkg/requester"
	"github.com/junminhong/member-center-service/pkg/responser"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	request := requester.Register{
		Email:    "test@gmail.com",
		Password: "test",
		NickName: "test",
	}
	/*mockMemberInfo := domain.MemberInfo{NickName: "test"}
	mockMember := domain.Member{
		Email:      "test@gmail.com",
		Password:   "test",
		MemberInfo: mockMemberInfo,
	}*/
	mockMemberUseCase := new(mocks.MemberUseCase)
	mockMemberUseCase.On("Register", mock.Anything, mock.AnythingOfType("*domain.Member")).Return(responser.Response{})
	handler := memberDeliver.MemberHandler{
		MemberUseCase: mockMemberUseCase,
	}
	r := gin.Default()
	r.POST("/api/v1/member", handler.Register)
	jsonString, _ := json.Marshal(request)
	log.Println(bytes.NewBuffer(jsonString))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/member", bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	log.Println(req)
	log.Println(w.Body)
}
