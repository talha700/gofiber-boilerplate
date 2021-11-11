package models


type SignIn struct {
	Username    string 
	Password string 
}

type SignUp struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Username string `json:"username"`
	Password string `json:"password" validate:"required,lte=255"`
	UserRole string `json:"-"`
}

type EditPassword struct {
	OldPass string  `json:"old_pass"`
	NewPass string	`json:"new_pass" validate:"required,lte=255"`
}

type LoginResponse struct {
	Error bool
	Msg string
	Token Tokens
}


type Tokens struct {
	Access string   `json:"access"`
	Refresh string  `json:"refresh"`
}



type ChangePass struct {
	Error bool `json:"error" example:"false"`
	Msg string 
}


type SignUpResponse struct {
	ChangePass
}

type RefreshTokenResponse struct {
	ChangePass
}
