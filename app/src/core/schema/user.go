package schema

type (
	UserSchema struct {
		ID       uint64 `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	UserCreateSchema struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
