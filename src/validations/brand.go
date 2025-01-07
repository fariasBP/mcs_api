package validations

type (
	CreateBrandParams struct {
		Name string `json:"name" validate:"required"`
	}
)
