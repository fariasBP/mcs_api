package validations

type (
	AddPermissionParams struct {
		IdUser     string `json:"id_user" validate:"required"`
		Permission string `json:"permission" validate:"required"`
	}
)
