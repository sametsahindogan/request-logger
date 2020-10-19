package request

import "time"

type GetByDomainRequestValidation struct {
	UserId  string
	Domain  string `validate:"required,url"`
	Uri     string `validate:"omitempty,uri"`
	Ip      string `validate:"omitempty,ip"`
	Created time.Time
	PerPage int64 `validate:"max=100"`
	Page    int64
	Sort    string
}
