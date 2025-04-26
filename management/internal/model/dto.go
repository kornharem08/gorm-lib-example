package model

type UserWithCompany struct {
	User
	CompanyName    string `json:"CompanyName"`
	CompanyAddress string `json:"CompanyAddress"`
}
