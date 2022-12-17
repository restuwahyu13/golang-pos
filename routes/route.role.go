package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/handlers"
	"github.com/restuwahyu13/golang-pos/middlewares"
	"github.com/restuwahyu13/golang-pos/repositories"
	"github.com/restuwahyu13/golang-pos/services"
)

func NewRouteRole(db *gorm.DB, router *gin.Engine) {
	repository := repositories.NewRepositoryRole(db)
	service := services.NewServiceRole(repository)
	handler := handlers.NewHandlerRole(service)

	route := router.Group("/api/v1/role")
	route.Use(middlewares.AuthToken())
	route.Use(middlewares.AuthRole(map[string]bool{"admin": true}))

	router.GET("/api/v1/role/ping", handler.HandlerPing)
	route.POST("/create", handler.HandlerCreate)
	route.GET("/results", handler.HandlerResults)
	route.GET("/result/:id", handler.HandlerResult)
	route.DELETE("/delete/:id", handler.HandlerDelete)
	route.PUT("/update:id", handler.HandlerUpdate)
}
