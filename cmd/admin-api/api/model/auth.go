package model

type (
	AuthRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	Token struct {
		Value string
	}
)
