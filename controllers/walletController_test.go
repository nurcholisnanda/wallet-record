package controllers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
	"github.com/nurcholisnanda/wallet-record/mocks"
	"github.com/nurcholisnanda/wallet-record/models"
	"github.com/nurcholisnanda/wallet-record/services"
	"github.com/stretchr/testify/assert"
)

var mockService = new(mocks.Service)

func TestNewController(t *testing.T) {
	type args struct {
		s services.Service
	}
	tests := []struct {
		name string
		args args
		want Controller
	}{
		{
			name: "Success test case",
			args: args{mockService},
			want: &controller{mockService},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_CreateRecord(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name string
		c    *controller
		args args
	}{
		{
			name: "Success test case",
			c:    &controller{mockService},
			args: args{&gin.Context{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Create a mock context and set json body
			jsonReq := `{"amount":1000, "datetime":"2019-10-05T13:00:00+00:00"}`
			req, _ := http.NewRequest("POST", "/records", strings.NewReader(jsonReq))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			tt.args.ctx, _ = gin.CreateTestContext(w)
			tt.args.ctx.Request = req

			mockService.On("CreateRecord", mock.Anything).Return(nil)

			//call CreateRecord
			tt.c.CreateRecord(tt.args.ctx)

			//Assert the response
			assert.Equal(t, http.StatusCreated, w.Code)
		})
	}
}

func Test_controller_GetHistory(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		c    *controller
		args args
	}{
		{
			name: "Success test case",
			c:    &controller{mockService},
			args: args{&gin.Context{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Create a mock context and set json body
			jsonReq := `{"startDatetime":"2019-10-05T13:00:00+00:00", "endDatetime":"2019-10-06T13:00:00+00:00"}`
			req, _ := http.NewRequest("POST", "/records/history", strings.NewReader(jsonReq))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			tt.args.ctx, _ = gin.CreateTestContext(w)
			tt.args.ctx.Request = req

			mockService.On("GetHistory", mock.Anything, mock.Anything).Return([]models.Record{}, nil)

			//call GetHistory
			tt.c.GetHistory(tt.args.ctx)

			//Assert the response
			assert.Equal(t, http.StatusOK, w.Code)
		})
	}
}

func Test_controller_GetLatest(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		c    *controller
		args args
	}{
		{
			name: "Success test case",
			c:    &controller{mockService},
			args: args{&gin.Context{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Create a mock context
			w := httptest.NewRecorder()
			tt.args.ctx, _ = gin.CreateTestContext(w)

			mockService.On("GetLatest", mock.Anything).Return(models.Record{}, nil)

			//call GetLatest
			tt.c.GetLatest(tt.args.ctx)

			//Assert the response
			assert.Equal(t, http.StatusOK, w.Code)
		})
	}
}
