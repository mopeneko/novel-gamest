package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/novel-gamest/api/interface/controller"
	"github.com/mopeneko/novel-gamest/api/interface/database"
)

var router *echo.Echo

func init() {
	router = echo.New()

	// JWTの秘密鍵をDBから取得
	secret := jwtSecret{}
	db.First(secret)

	userController := controller.UserController{
		UserRepository:    database.UserRepository{DB: db},
		UserTokenProvider: NewJWTProvider(secret.secret),
	}
	router.POST("/sign_in", userController.SignIn)
	router.POST("/sign_up", userController.SignUp)
	router.GET("/users/:id", userController.GetByID)
}

// Run HTTP server
func Run() {
	router.Logger.Fatal(router.Start(":1323"))
}
