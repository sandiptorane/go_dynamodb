package main

import (
	"github.com/gin-gonic/gin"
	"go_dynamodb/handler"
	"go_dynamodb/pkg/response"
)

// RegisterRoutes registers routes and handlers and return router
func RegisterRoutes(app *handler.Application) *gin.Engine {
	r := gin.Default()

	r.NoRoute(NoRoute)

	r.POST("/article", app.SaveArticle)
	r.PUT("/article", app.UpdateArticle)
	r.GET("/article/:title/:author", app.GetArticle)
	r.GET("/article", app.GetAllArticles)
	r.DELETE("/article", app.DeleteArticle)

	r.GET("/health", HealthCheck)

	return r
}

func NoRoute(c *gin.Context) {
	response.NotFound(c, "route not found", nil)
}

func HealthCheck(c *gin.Context) {
	response.Success(c, "ok", nil)
}
