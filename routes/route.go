package routes

import (
	"MyGram/controllers/photos-controllers/create"
	"MyGram/controllers/user-controllers/login"
	"MyGram/controllers/user-controllers/register"
	handlerCreate "MyGram/handlers/photos-handler/create"
	handlerLogin "MyGram/handlers/users-handler/login"
	register2 "MyGram/handlers/users-handler/register"
	"MyGram/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoute(db *gorm.DB, route *gin.Engine) {
	registerRepository := register.NewRepositoryRegister(db)
	registerService := register.NewServiceRegister(registerRepository)
	registerHandler := register2.NewHandlerRegister(registerService)

	loginRepository := login.NewLoginrepository(db)
	loginService := login.NewLoginService(loginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	groupRoute := route.Group("/users")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
}
func PhotosRoute(db *gorm.DB, route *gin.Engine) {

	createRepository := create.NewCreateRepository(db)
	createService := create.NewCreatePhotoService(createRepository)
	createHandler := handlerCreate.NewHandlerCreatePhotos(createService)

	groupRoute := route.Group("/photos").Use(middleware.Auth())
	groupRoute.POST("/", createHandler.CreateStudentHandler)
}
