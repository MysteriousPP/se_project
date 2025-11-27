package userin

type UserCreateInp struct {
	Phone    string `json:"phone" v:"required"`
	Password string `json:"password" v:"required"`
	Nickname string `json:"nickname" v:"required"`
	Role     int8   `json:"role" v:"required"`
}

type UserCreateModel struct {
	UserId int    `json:"userId"`
	Token  string `json:"token"`
}

type UserLoginInp struct {
	Phone    string `json:"phone" v:"required"`
	Password string `json:"password" v:"required"`
}

type UserLoginModel struct {
	UserId int    `json:"userId"`
	Token  string `json:"token"`
}

type BindElderInp struct {
	ChildUserId int    `json:"childUserId" v:"required"`
	ElderUserId int    `json:"elderUserId" v:"required"`
	BindCode    string `json:"bindCode" v:"required"`
}

type BindElderModel struct {
}
