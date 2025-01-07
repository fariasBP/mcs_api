package validations

type (
	LoginParams struct {
		Identifier string `json:"user" validate:"required"`
		Pwd        string `json:"pwd" validate:"required"`
	}
	SignUpParams struct {
		IdName string `json:"id_name" validate:"required,lowercase,min=5,max=30"`
		Name   string `json:"name" validate:"required,lowercase,min=3,max=30"`
		Lname  string `json:"lname" validate:"required,lowercase,min=3,max=30"`
		Email  string `json:"email" validate:"required,email"`
		Pwd    string `json:"pwd" validate:"required"`
		Bth    string `json:"bth" validate:"required,date"`
	}
)
