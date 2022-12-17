package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/handlers"
	"github.com/restuwahyu13/golang-pos/middlewares"
	"github.com/restuwahyu13/golang-pos/repositories"
	"github.com/restuwahyu13/golang-pos/services"
)

func NewRouteCustomer(db *gorm.DB, router *gin.Engine) {
	repository := repositories.NewRepositoryCustomer(db)
	service := services.NewServiceCustomer(repository)
	handler := handlers.NewHandlerCustomer(service)

	route := router.Group("/api/v1/customer")
	route.Use(middlewares.AuthToken())
	route.Use(middlewares.AuthRole(map[string]bool{"admin": true, "outlet": true}))

	router.GET("/api/v1/customer/ping", handler.HandlerPing)
	route.POST("/create", handler.HandlerCreate)
	route.GET("/results", handler.HandlerResults)
	route.GET("/result/:id", handler.HandlerResult)
	route.DELETE("/delete/:id", handler.HandlerDelete)
	route.PUT("/update:id", handler.HandlerUpdate)
}
