package authdto

type SignUpRequest struct {
	Name     string `json:"name" validate:"required" message:"name is required"`
	Email    string `json:"email" validate:"email" message:"email is invalid"`
	Password string `json:"password" validate:"required|min_len:8" message:"password must be at least 8 characters long"`
	Role     string `json:"role" validate:"required|in:admin,user" message:"role must be either admin or user"`
}
