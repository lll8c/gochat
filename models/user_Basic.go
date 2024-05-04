package models

import (
	"fmt"
	"gochat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d)"`
	Email         string `valid:"email"`
	Avatar        string //头像
	Identity      string //token字符串
	ClientIP      string
	ClientPort    string
	Salt          string    //MD5加密随机盐值
	LoginTime     time.Time //登录时间
	HeartbeatTime time.Time //心跳时间
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"` //退出时间
	gorm.Model
}

func (UserBasic) TableName() string {
	return "user_basics"
}

// GetUserList 获取所有user
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

// CreateUser 添加用户
func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func FindUserByNameAndPassword(name, password string) (user UserBasic) {
	utils.DB.Where("name = ? and pass_word = ?", name, password).First(&user)
	return user
}

func FindUserByName(name string) (user UserBasic) {
	utils.DB.Where("name = ?", name).Find(&user)
	return
}

// UpdateToken 生成并保存token
func UpdateToken(user *UserBasic) {
	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.Md5Encode(str)
	utils.DB.Model(user).Where("id = ?", user.ID).Update("identity", temp)
}

func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("phone = ?", phone).First(&user)
}
func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, Password: user.Password, Phone: user.Phone, Email: user.Email})
}

// FindByID 查找某个用户
func FindByID(id uint) UserBasic {
	user := UserBasic{}
	utils.DB.Where("id = ?", id).First(&user)
	return user
}
