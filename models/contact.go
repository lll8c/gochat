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

// AddFriend 添加好友
func AddFriend(userId uint, targetId uint) (int, string) {
	user := UserBasic{}
	user = FindByID(targetId)
	if user.ID == 0 {
		return -1, "用户不存在"
	}
	//判断是否已为好友
	temp := FindContact(userId, targetId)
	if temp.ID != 0 {
		return -1, "好友已存在"
	}
	//开启事务添加两条记录
	tx := utils.DB.Begin()
	//事务一旦开始，不论什么异常都会RollBackx
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var contact, contact2 Contact
	contact.OwnerId = userId
	contact.TargetId = targetId
	contact.Type = 1
	if err := tx.Create(&contact).Error; err != nil {
		tx.Rollback()
		return -1, "添加好友失败"
	}
	contact2.OwnerId = targetId
	contact2.TargetId = userId
	contact2.Type = 1
	if err := tx.Create(&contact2).Error; err != nil {
		tx.Rollback()
		return -1, "添加好友失败"
	}
	tx.Commit()
	return 0, "添加好友成功" //成功
}

func FindContact(userID uint, targetID uint) Contact {
	var contact Contact
	utils.DB.Where("owner_id=? and target_id=? and type = 1", userID, targetID).First(&contact)
	return contact
}
