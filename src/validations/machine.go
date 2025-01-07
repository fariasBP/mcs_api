package validations

type (
	CreateMachineParams struct {
		CompanyId     string `json:"company_id" validate:"required,lowercase,min=3,max=30"`
		MachineTypeId string `json:"machine_type_id" validate:"required"`
		BrandId       string `json:"brand_id" validate:"required"`
		Serial        string `json:"serial" validate:"required,min=3,max=30"`
		Model         string `json:"model" validate:"required,min=3,max=30"`
	}
)
