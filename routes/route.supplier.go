package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/handlers"
	"github.com/restuwahyu13/golang-pos/middlewares"
	"github.com/restuwahyu13/golang-pos/repositorys"
	"github.com/restuwahyu13/golang-pos/services"
)

func NewRouteSupplier(db *gorm.DB, router *gin.Engine) {
	repository := repositorys.NewRepositorySupplier(db)
	service := services.NewServiceSupplier(repository)
	handler := handlers.NewHandlerSupplier(service)

	route := router.Group("/api/v1/supplier")
	route.Use(middlewares.AuthToken())
	route.Use(middlewares.AuthRole(map[string]bool{"admin": true, "supplier": true}))

	router.GET("/api/v1/supplier/ping", handler.HandlerPing)
	route.POST("/create", handler.HandlerCreate)
	route.GET("/results", handler.HandlerResults)
	route.GET("/result/:id", handler.HandlerResult)
	route.DELETE("/delete/:id", handler.HandlerDelete)
	route.PUT("/update:id", handler.HandlerUpdate)
}
