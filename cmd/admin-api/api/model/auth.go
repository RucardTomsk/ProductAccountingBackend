package model

type (
	AuthRequest struct {
		Email    string
		Password string
	}

	RegisterRequest struct {
		Email    string
		Password string
		Role     string
	}

	Token struct {
		Value string
	}
)
