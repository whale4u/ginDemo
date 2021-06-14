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

// 注：存在bug，无论是否存在都提示删除成功！！！

func (passwd *Passwd) DeletePassword(id int64) (Result Passwd) {
	passwd.Id = id
	//if err := orm.DB.Where("id = ?", id).Delete(&passwd).Error; err != nil {
	//	return
	//}
	if err := orm.DB.Delete(&passwd).Error; err != nil {
		return
	}
	Result = *passwd
	return
}
