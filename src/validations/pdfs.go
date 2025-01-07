package validations

type (
	GenerateBasicPdfsParams struct {
		ServiceId string `json:"service_id" validate:"required"`
	}
)
