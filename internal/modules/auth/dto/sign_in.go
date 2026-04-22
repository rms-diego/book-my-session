package authdto

type SignInRequest struct {
	Email    string `json:"email" validate:"email" message:"email is invalid"`
	Password string `json:"password" validate:"required|min_len:8" message:"password must be at least 8 characters long"`
}
