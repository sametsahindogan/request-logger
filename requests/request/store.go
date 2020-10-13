package request

type StoreProcessRequestValidation struct {
	UserId    string `json:"user_id" validate:"required,min=3,max=40"`
	IpAddress string `json:"ip_address" validate:"required,ip"`
	Uri       string `json:"uri" validate:"required,uri"`
	Domain    string `json:"domain" validate:"required,url"`
}