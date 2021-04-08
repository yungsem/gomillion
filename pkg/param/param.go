package param

type JobParam interface {
	String() string
}

type EqpStatus struct {
	Code string `json:"code"`
	Status string `json:"status"`
}

func (e *EqpStatus) String() string {
	return e.Code + ";" + e.Status
}
