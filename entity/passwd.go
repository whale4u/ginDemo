package entity

import (
	"fmt"
	orm "ginDemo/common"
)

type Passwd struct {
	Id       int64
	Type     string `json:"type"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
	Note     string `json:"note"`
}

func (passwd Passwd) InsertPasswd() (Result bool, Id int64) {
	result := orm.DB.Create(&passwd)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error)
		return false, -1
	}
	return true, passwd.Id
}
