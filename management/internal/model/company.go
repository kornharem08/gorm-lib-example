package model

type Company struct {
	Code    string `json:"Code" gorm:"column:Code;primaryKey"`
	Name    string `json:"Name" gorm:"column:Name"`
	Address string `json:"Address" gorm:"column:Address"`
}

func (Company) TableName() string {
	return "companies"
}
