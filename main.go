package main

import (
	//"net/http"

	//"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"

	"github.com/sac-production-2019/sac-go-backend/route"
	model "github.com/sac-production-2019/sac-go-backend/model"

	//"github.com/sac-production-2019/sac-go-backend/controller"
)


func main() {
	/*
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	

	e.GET("/",func(c echo.Context) error{
		return c.String(http.StatusOK,"Hello World!\n")
	})

	e.POST("/User",controller.CreateUser)
	e.GET("/User/:id",controller.GetUser)
	*/
	model.InitUsers()
	e := route.InitRouter()
	e.Logger.Fatal(e.Start(":1323"))
}







//regexMatch(r.act, p.act)