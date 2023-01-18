package routes

import (
	"reflect"
	"testing"

	"github.com/nurcholisnanda/wallet-record/controllers"
	"github.com/nurcholisnanda/wallet-record/mocks"
)

var mockCtrl = new(mocks.Controller)

func TestNewRouter(t *testing.T) {
	type args struct {
		c controllers.Controller
	}
	tests := []struct {
		name string
		args args
		want Router
	}{
		{
			name: "Success test case",
			args: args{
				mockCtrl,
			},
			want: &router{mockCtrl},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_router_SetupRouter(t *testing.T) {
	tests := []struct {
		name string
		r    *router
	}{
		{
			name: "Success test case",
			r:    &router{mockCtrl},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.r.SetupRouter()
			if len(e.Routes())-1 != 3 {
				t.Errorf("Expected 3 routes to be registered, got %d", len(e.Routes()))
			}
			routePaths := []string{"GET /records/latest", "POST /records", "POST /records/history"}
			for _, routePath := range routePaths {
				routeExists := false
				for _, route := range e.Routes() {
					if routePath == route.Method+" "+route.Path {
						routeExists = true
						break
					}
				}
				if !routeExists {
					t.Errorf("Route %s is not registered", routePath)
				}
			}
		})
	}
}
