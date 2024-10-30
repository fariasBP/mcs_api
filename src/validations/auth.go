package validations

type (
	LoginParams struct {
		Identifier string `json:"user" validate:"required"`
		Pwd        string `json:"pwd" validate:"required"`
	}
)
