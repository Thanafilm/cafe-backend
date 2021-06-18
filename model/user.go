package model

type User struct {
	Id       uint
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Username string `json:"username"`
	Password []byte `json:"-"`
	MenuID   string `json:"menuid"`
	Email    string `json:"email" gorm:"unique"`
}
