package validations

type (
	CreateProtocolParams struct {
		Acronym     string `json:"acronym" validate:"required"`
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
)
