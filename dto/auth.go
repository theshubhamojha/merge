package dto

type LoginRequest struct {
	EmailID  string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type AccountDetails struct {
}
