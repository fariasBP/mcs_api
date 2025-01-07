package validations

type (
	CompanyParams struct {
		Name        string `json:"name" validate:"required,lowercase,min=2,max=30"`
		Manager     string `json:"manager" validate:"required,lowercase,min=3,max=30"`
		Location    string `json:"location" validate:"required,location"`
		Description string `json:"description" validate:"required"`
	}
)
