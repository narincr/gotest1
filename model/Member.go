package model

type MemberGetId struct {
	Id int `json:"Id"`
}

type MemberGetList struct {
	FirstName string `gorm:"column:FirstName" json:"FirstName"`
	LastName  string `gorm:"column:LastName" json:"LastName"`
	Phone     string `gorm:"column:Phone" json:"Phone"`
}

type MemberInfo struct {
	MemberID  int32  `gorm:"column:MemberID;primaryKey;autoIncrement:true" json:"MemberID"`
	FirstName string `gorm:"column:FirstName" json:"FirstName"`
	LastName  string `gorm:"column:LastName" json:"LastName"`
	Email     string `gorm:"column:Email" json:"Email"`
	Phone     string `gorm:"column:Phone" json:"Phone"`
	BirthDate string `gorm:"column:BirthDate" json:"BirthDate"`
	Address   string `gorm:"column:Address" json:"Address"`
	City      string `gorm:"column:City" json:"City"`
	State     string `gorm:"column:State" json:"State"`
	ZipCode   string `gorm:"column:ZipCode" json:"ZipCode"`
}
