package userController

import (
	"github.com/labstack/echo"
	"net/http"
	//"github.com/labstack/echo/middleware"
	model "github.com/sac-production-2019/sac-go-backend/model"
	"fmt"
	"strconv"
	"errors"
	"github.com/sac-production-2019/sac-go-backend/utils/gwt"
	"github.com/sac-production-2019/sac-go-backend/utils"
)



func CreateUser(c echo.Context) error {
	u := &model.User{ID:model.Seq,}
	if err := c.Bind(u); err !=nil{
		return err
	}
	model.Users[model.Seq] = u
	model.Seq++
	fmt.Println(model.Users)
	return c.JSON(http.StatusCreated,u)
}
//path /user/:id
func GetUser(c echo.Context) error{
	fmt.Println("heeloasdasdasd")
	id,_ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK,model.Users[id])
}

func Login(c echo.Context) error{
	//role := c.FormValue("role")
	id,_ := strconv.Atoi(c.Param("id"))
	if model.Users[id]!=nil{
		token, err := gwt.GenerateToken(model.Users[id])
		if err !=nil{
			c.JSON(utils.Error(err.Error()))
			return err
		}
		
		c.JSON(utils.Success("success",utils.LoginResult{token,*model.Users[id]}))
		return nil
	}
	return errors.New("not found user")
}


func Hello(c echo.Context) error{
	return c.JSON(utils.Success("hello world"))
}