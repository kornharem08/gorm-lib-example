package model

type User struct {
	Contactid      string `json:"Contactid" gorm:"column:Contactid"`
	ADUserId       string `json:"ADUserId" gorm:"column:ADUserId"`
	EmployeeId     string `json:"EmployeeId" gorm:"column:EmployeeId"`
	FirstName      string `json:"FirstName" gorm:"column:FirstName"`
	LastName       string `json:"LastName" gorm:"column:LastName"`
	LocalFirstName string `json:"LocalFirstName" gorm:"column:LocalFirstName"`
	LocalLastName  string `json:"LocalLastName" gorm:"column:LocalLastName"`
	EmployeeTag    string `json:"EmployeeTag" gorm:"column:EmployeeTag"`
	CompanyCode    string `json:"CompanyCode" gorm:"column:CompanyCode"`
}

func (User) TableName() string {
	return "users" // Replace with your desired table name, e.g., "user_table"
}
