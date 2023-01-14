package model

type UserModel struct {
	Username    string   `json:"username"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type AuthToken struct {
	Token string `json:"Token"`
	UserModel
}

type JwtBody struct {
	Username    string   `json:"username"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type RegisterUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=12,regexp=^.*[a-zA-Z].*$"`
}

type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=12,regexp=^.*[a-zA-Z].*$"`
}
