// routes.go

package main

import (
	"iad/api"
	// "iad/docs"
	"iad/middleware"
	"iad/page"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func initializeRoutes() {
	router.Static("/static/css", "./static/css")
	router.Static("/static/img", "./static/img")
	router.Static("/static/js", "./static/js")

	// Swagger docs
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.GET("/", page.RedirectIndexPage)

	healthRoutes := router.Group("/health", middleware.CORSMiddleware())
	{
		healthRoutes.GET("/healthy", api.Health)
		healthRoutes.GET("/ready", api.Readiness)
	}

	apiRoutes := router.Group("/api", middleware.CORSMiddleware())
	{
		v1Routes := apiRoutes.Group("/v1")
		{
			v1Routes.POST("/aws", api.AWSPost, middleware.CORSMiddleware())
		}
	}

	// uiRoutes := router.Group("/ui", middleware.CORSMiddleware())
	// {
	// 	// Add UI routes here
	// }

	// htmxRoutes := router.Group("/htmx", middleware.CORSMiddleware(), middleware.EnsureLoggedInAPI())
	// {
	// 	// Add HTMX routes here
	// }
}
