package model

type (
	User struct{
		ID int `json:"id"`
		Name string `json:"name"`
		Role string `json:"role"`
	}
)

var (
	Users = map[int]*User{}
	Seq =1

)

func InitUsers(){
	Users[1] = &User{
		1,
		"zzt",
		"admin",
	}

	Users[2] = &User{
		2,
		"zzt1",
		"member",
	
	}
}




