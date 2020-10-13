package request

import "time"

type GetByDomainRequestValidation struct {
	Domain  string `validate:"required,url"`
	Created time.Time
	UserId  string
	Limit   int64 `validate:"max=100"`
	Offset  int64
	Sort    string
}
