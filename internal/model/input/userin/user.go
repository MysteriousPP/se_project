package userin

type UserCreateInp struct {
	Nickname string `json:"username" v:"required"`
	Password string `json:"password" v:"required"`
	Role     int8   `json:"role" v:"required"`
	Phone    string `json:"phone" v:"required"`
	Email    string `json:"email"`
}

type UserCreateModel struct {
	UserId   int    `json:"userId"`
	Nickname string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type UserLoginInp struct {
	Phone    string `json:"phone" v:"required"`
	Username string `json:"username"`
	Password string `json:"password" v:"required"`
}

type UserLoginModel struct {
	Token    string `json:"token"`
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type BindElderInp struct {
	ChildUserId int    `json:"childUserId" v:"required"`
	ElderUserId int    `json:"elderUserId" v:"required"`
	Relation    string `json:"relation"`
	BindCode    string `json:"bindCode" v:"required"`
}

type BindElderModel struct {
	BindId   int    `json:"bindId"`
	ChildId  int    `json:"childId"`
	ElderId  int    `json:"elderId"`
	Relation string `json:"relation"`
}
