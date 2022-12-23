package echo

type JSONLogin struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// swagger
type JSONResult struct {
	Message string `json:"message"`
}

type JSONToken struct {
	Token string `json:"token"`
}
