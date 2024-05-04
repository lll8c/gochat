package models

import (
	"gochat/utils"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name    string
	OwnerId uint
	Img     string
	Desc    string
}

// CreateCommunity 创建群
func CreateCommunity(community Community) (int, string) {
	tx := utils.DB.Begin()
	if len(community.Name) == 0 {
		return -1, "群名称不能为空"
	}
	if community.OwnerId == 0 {
		return -1, "请先登录"
	}
	if err := tx.Create(&community).Error; err != nil {
		tx.Rollback()
		return -1, "建群失败"
	}

	//群主加入该群，添加群关系
	contact := Contact{
		OwnerId:  community.OwnerId,
		TargetId: community.ID,
		Type:     2,
	}
	if err := tx.Create(&contact).Error; err != nil {
		tx.Rollback()
		return -1, "添加群关系失败"
	}
	tx.Commit()
	return 0, "建群成功"
}

// SearchCommunity 查找该用户加入的所有群聊
func SearchCommunity(userId uint) []*Community {
	contacts := []*Contact{}
	//先找到该用户加入的所有群聊的id
	ids := []uint{}
	utils.DB.Where("owner_Id = ?", userId).Find(&contacts)
	for _, v := range contacts {
		ids = append(ids, v.TargetId)
	}
	//查找群聊
	data := make([]*Community, 0)
	utils.DB.Where("id in ?", ids).Find(&data)
	return data
}

// JoinGroup 加入群聊
func JoinGroup(userId uint, comId uint) (int, string) {
	contact := Contact{
		OwnerId:  userId,
		TargetId: comId,
		Type:     2,
	}
	community := Community{}
	utils.DB.Where("id=? or name=?", comId, comId).Find(&community)
	if community.Name == "" {
		return -1, "没有找到群"
	}
	utils.DB.Where("owner_id=? and target_id=? and type = 2 ", userId, comId).Find(&contact)
	if !contact.CreatedAt.IsZero() {
		return -1, "已加过此群"
	} else {
		contact.TargetId = community.ID
		utils.DB.Create(&contact)
		return 0, "加群成功"
	}
}

// FindContactByGroupId 查找群聊成员
func FindContactByGroupId(groupId uint) []*Contact {
	data := []*Contact{}
	utils.DB.Where("target_id = ? and type = 2", groupId).Find(&data)
	return data
}
