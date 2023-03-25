package dto

type RequestRegisterDTO struct {
	PhoneNumber string `json:"phoneNumber"`
	Role        string `json:"role"`
}

type LoginDTO struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type CheckCodeRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Code        string `json:"code"`
}
