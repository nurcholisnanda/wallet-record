package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nurcholisnanda/wallet-record/controllers"
	"github.com/nurcholisnanda/wallet-record/docs"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// server is a global variable that is set to a gin router with the default middleware.
var server = gin.Default()

// applicationBasePath is the base path for all routes in the application.
const applicationBasePath = "/"

type Router interface {
	SetupRouter() *gin.Engine
}

type router struct {
	controller controllers.Controller
}

// NewRouter is a constructor function that returns an instance of the router struct with the provided controller.
func NewRouter(c controllers.Controller) Router {
	return &router{
		controller: c,
	}
}

// SetupRouter ... Configure routes
// This function sets up the routes for the application, including the routes for the swagger documentation.
// It also sets up group routes for the applicationBasePath and assigns handlers for each route.
func (r *router) SetupRouter() *gin.Engine {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Wallet Balance API"
	docs.SwaggerInfo.Description = "API for Insert, Update and Get Balance of your wallet"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "nimble-monument-374407.an.r.appspot.com"
	docs.SwaggerInfo.BasePath = applicationBasePath
	docs.SwaggerInfo.Schemes = []string{"https"}

	// Group route for the application base path
	grp1 := server.Group(applicationBasePath)
	{
		// Route for getting the latest record
		grp1.GET("records/latest", func(ctx *gin.Context) {
			r.controller.GetLatest(ctx)
		})
		// Route for creating a new record
		grp1.POST("records", func(ctx *gin.Context) {
			r.controller.CreateRecord(ctx)
		})
		// Route for getting history of records
		grp1.POST("records/history", func(ctx *gin.Context) {
			r.controller.GetHistory(ctx)
		})
	}

	// Route for serving the swagger documentation
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return server
}
