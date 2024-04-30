package models

import (
	"fmt"
	"gochat/utils"

	"gorm.io/gorm"
)

// Contact 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的谁 /群 ID
	Type     int  //对应的类型  1好友  2群  3xx
	Desc     string
}

func (table *Contact) TableName() string {
	return "contacts"
}

// SearchFriend 查找这个用户的所有好友
func SearchFriend(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	ids := []uint{}
	users := []UserBasic{}
	utils.DB.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	for _, v := range contacts {
		fmt.Println(v)
		ids = append(ids, v.TargetId)
	}
	utils.DB.Where("id in ?", ids).Find(&users)
	return users
}
