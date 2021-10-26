package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/handlers"
	"github.com/restuwahyu13/golang-pos/repositorys"
	"github.com/restuwahyu13/golang-pos/services"
)

func NewRouteUser(db *gorm.DB, router *gin.Engine) {
	repositoryUser := repositorys.NewRepositoryUser(db)
	serviceUser := services.NewServiceUser(repositoryUser)
	handlerUser := handlers.NewHandlerUser(serviceUser)

	route := router.Group("/api/v1/auth")

	route.GET("/ping", handlerUser.HandlerPing)
	route.POST("/register", handlerUser.HandlerRegister)
	route.POST("/login", handlerUser.HandlerLogin)
}
