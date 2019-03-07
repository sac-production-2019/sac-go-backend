package auth

import(
	//"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/casbin/casbin"
	"log"
	"github.com/sac-production-2019/sac-go-backend/utils"
	"fmt"
	//"github.com/sac-production-2019/sac-go-backend/"
)

var(
	conf = "config/auth_model.conf"  //"../config/auth_model.conf"
	policy = "config/policy.csv" //../config/policy.csv"
)

//这里的返回error为nil，而在jwt里面返回了error，不一样，运行的时候看一下
func AuthCheckRole(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error{
		fmt.Println("hello1")
		claims := c.Get("claims").(*CustomClaims)
		//name := claims.Name
		role := claims.Role
		enf,err := casbin.NewEnforcerSafe(conf,policy)
		
		if err !=nil{
			log.Fatal(err)
		}
		
		fmt.Println(role)
		fmt.Println(c.Request().URL.Path)
		fmt.Println(c.Request().Method)

		res := enf.Enforce(role,c.Request().URL.Path,c.Request().Method)
		
		/*
		if err1 != nil {
			fmt.Println("enforce failed")
			fmt.Println(err1)
			fmt.Println(res)
			c.JSON(utils.Error("服务器错误"))
		}
		*/

		if res{
			fmt.Println("code comes here")
			return next(c) 
		} else{
			c.JSON(utils.UnAuth("您没有此权限"))
			return nil
		}
	}
}

