package auth

type AuthSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSignInResponse struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}
