package auth

type SignUp struct {
	UserID uint   `json:"userId"`
	Role   string `json:"role"`
}
