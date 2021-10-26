package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/handlers"
	"github.com/restuwahyu13/golang-pos/middlewares"
	"github.com/restuwahyu13/golang-pos/repositorys"
	"github.com/restuwahyu13/golang-pos/services"
)

func NewRouteOutlet(db *gorm.DB, router *gin.Engine) {
	repository := repositorys.NewRepositoryOutlet(db)
	service := services.NewServiceOutlet(repository)
	handler := handlers.NewHandlerOutlet(service)

	route := router.Group("/api/v1/outlet")
	route.Use(middlewares.AuthToken())
	route.Use(middlewares.AuthRole(map[string]bool{"admin": true, "merchant": true, "outlet": true}))

	router.GET("/api/v1/outlet/ping", handler.HandlerPing)
	route.POST("/create", handler.HandlerCreate)
	route.GET("/results", handler.HandlerResults)
	route.GET("/result/:id", handler.HandlerResult)
	route.DELETE("/delete/:id", handler.HandlerDelete)
	route.PUT("/update:id", handler.HandlerUpdate)
}
