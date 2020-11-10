package viewmodel

type SignInUserRequest struct {
	Email    string `json:"login"`
	Password string `json:"password"`
}
