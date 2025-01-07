package validations

type (
	ProblemParams struct {
		Problem  string `json:"problem" validate:"required"`
		Solution string `json:"solution" validate:"required"`
	}
	MaterialParams struct {
		Name   string `json:"name" validate:"required"`
		Number int    `json:"number" validate:"required"`
		Price  int    `json:"price" validate:"required"`
	}
	ProtocolParams struct {
		ProtocolId string          `json:"protocol_id" validate:"required"`
		Status     int             `json:"status" validate:"required"`
		Note       string          `json:"note" validate:"required"`
		Problems   []ProblemParams `json:"problems" validate:"required"`
	}
	CreateServiceParams struct {
		MachineId string           `json:"machine_id" validate:"required"`
		Comments  string           `json:"comments" validate:"required"`
		StartedAt string           `json:"started_at" validate:"required"`
		EndedAt   string           `json:"ended_at" validate:"required"`
		Status    int              `json:"status" validate:"required"`
		Materials []MaterialParams `json:"materials" validate:"required"`
		Protocols []ProtocolParams `json:"protocols" validate:"required"`
	}
)
