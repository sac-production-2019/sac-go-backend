package route

import(
	"github.com/labstack/echo"
	"github.com/sac-production-2019/sac-go-backend/authorization"
	//"net/http"
	"github.com/sac-production-2019/sac-go-backend/controller"
	"github.com/labstack/echo/middleware"
)

func InitRouter() *echo.Echo{
	e := echo.New()
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.GET("/user/:id",userController.GetUser)
	e.GET("/login/users/:id",userController.Login)
	e.POST("/register",userController.CreateUser)

	authRequired := e.Group("/apis")
	authRequired.Use(auth.JWTAuth())
	authRequired.Use(auth.AuthCheckRole)
	
		authRequired.GET("/user/:id",userController.GetUser)	
	
	
		


	return e
}


