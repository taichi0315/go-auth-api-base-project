package auth

type AuthSignUpRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSignUpResponse struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}
