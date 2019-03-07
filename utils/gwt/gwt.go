package gwt


import(
	"github.com/sac-production-2019/sac-go-backend/model"
	//"github.com/labstack/echo"
	"github.com/sac-production-2019/sac-go-backend/authorization"
	jwtgo "github.com/dgrijalva/jwt-go"
	//"github.com/sac-production-2019/sac-go-backend/utils"
	"time"
)




func GenerateToken(u *model.User) (string,error){
	k := auth.NewKey()
	claims := auth.CustomClaims{
		u.Name,
		u.Role,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), //签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), //签名过期时间 一小时
			Issuer:    "sysu-sac",                        //签名发行者
		},

	}
	token,err := k.CreateToken(claims)
	/*
	if err != nil{
		e.JSON(utils.Error(err.Error()))
	}
	*/
	if err != nil{
		return "",err
	}
	return token,nil
}


