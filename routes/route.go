package routes

import (
	"MyGram/controllers/user-controllers/register"
	register2 "MyGram/handlers/users-handler/register"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoute(db *gorm.DB, route *gin.Engine) {
	registerRepository := register.NewRepositoryRegister(db)
	registerService := register.NewServiceRegister(registerRepository)
	registerHandler := register2.NewHandlerRegister(registerService)

	groupRoute := route.Group("/users")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
}
