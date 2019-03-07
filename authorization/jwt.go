package auth

import(
	"github.com/dgrijalva/jwt-go"
	"errors"
	"github.com/labstack/echo"
	"github.com/sac-production-2019/sac-go-backend/utils"
	"fmt"
)

type Key struct{
	SigningKey []byte
}


var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "sac-sysu"
)

type CustomClaims struct{
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.StandardClaims
}


func SetSignKey(key string) bool{
	SignKey = key
	return true
}

func GetSignKey() string{
	return SignKey

}

func NewKey() *Key{
	return &Key{
		[]byte(GetSignKey()),
	}
}


func (k *Key) 	CreateToken(claims CustomClaims) (string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(k.SigningKey)
}

func (k * Key) ParseToken(tokenString string )(*CustomClaims, error){
	token,err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token)(interface{}, error){
		return k.SigningKey, nil
	})
	fmt.Println("token")
	fmt.Println(token)
	if err!= nil{
		if ve, ok := err.(*jwt.ValidationError);ok{
			if ve.Errors&jwt.ValidationErrorMalformed !=0{
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired !=0{
				return nil, TokenExpired
			} else if ve.Errors& jwt.ValidationErrorNotValidYet !=0{
				return nil,TokenNotValidYet
			} else{
				return nil, TokenInvalid
			}			
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok&& token.Valid{
		return claims,nil
	}
	return nil, TokenInvalid
}


func JWTAuth() echo.MiddlewareFunc{

	return func (next echo.HandlerFunc) echo.HandlerFunc{
		return func(ec echo.Context) error {
			//token := e.Request().Header.GET("")
			var errReturn error
			tokenString := ec.FormValue("Token")
			if tokenString ==""{
				//header 查找token
				tokenString = ec.Request().Header.Get("Token")//(echo.HeaderAuthorization)
				
				if tokenString ==""{
					ec.JSON(utils.ErrJWT(`请重新登录`,`未发现JWT`))
					errReturn = errors.New("未发现JWT")
					return errReturn
				}
				//tokenString = tokenString[7:]
			}
			k := NewKey()
			claims,err := k.ParseToken(tokenString)
			if err!=nil{
				fmt.Println(err)
				if err == TokenExpired{
					ec.JSON(utils.ErrJWT("请重新登录","JWT过期"))
					errReturn = errors.New("JWT Expired")
					return errReturn
					
				}
				ec.JSON(utils.ErrJWT("请重新登录",err))
				errReturn = errors.New("JWT验证失败")
				return errReturn
			}
			//ec.Response().Header().Set(echo.HeaderServer,"devfuck")
			ec.Set("claims",claims)
			return next(ec)
		}
	}

}

