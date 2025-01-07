package validations

type (
	CreateMachineTypeParams struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
)
