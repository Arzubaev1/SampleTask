package models

type PhoneNumberPrimaryKey struct {
	Id string `json:"id"`
}

type CreatePhoneNumber struct {
	Userid      string `json:"user_id"`
	Phone       string `json:"phone"`
	Isfax       bool   `json:"is_fax"`
	Description string `json:"description"`
}
type PhoneNumber struct {
	Id          string `json:"id"`
	UserId      string `json:"user_id"`
	Phone       string `json:"phone"`
	Isfax       bool   `json:"is_fax"`
	Description string `json:"description"`
}
type UpdatePhoneNumber struct {
	Id          string `json:"id"`
	Userid      string `json:"user_id"`
	Phone       string `json:"phone"`
	Isfax       bool   `json:"is_fax"`
	Description string `json:"description"`
}
type PhoneNumberGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type PhoneNumberGetListResponse struct {
	Count        int            `json:"count"`
	PhoneNumbers []*PhoneNumber `json:"phone_numbers"`
}
