package models

import (
	"code/utils"
	"fmt"

	"gorm.io/gorm"
)

// 用户结构体
type User struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LoginTime     *utils.LocalTime
	HeartbeatTime *utils.LocalTime
	LoginOutTime  *utils.LocalTime `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

//生成User表
func (table *User) TableName() string {
	return "t_user"
}

//获取用户列表
func GetUserList() []*User {
	data := make([]*User, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

//创建用户
func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}

//删除用户
func DeleteUser(user User) *gorm.DB {
	return utils.DB.Delete(&user)
}

//更新用户信息
// func UpdateUser(user User) *gorm.DB {
// 	fmt.Println(user.Name)
// 	return utils.DB.Model(&user).Updates(
// 		map[string]interface{}{
// 			"name":     user.Name,
// 			"password": user.Password,
// 		},
// 		// User{Name: user.Name, Password: user.Password}
// 	)
// }

func UpdateUser(user User) *gorm.DB {
	fmt.Println(user.Name)
	return utils.DB.Model(&user).Updates(User{Name: user.Name, Password: user.Password})
}

func UpdateInfo(user User) bool {
	if err := utils.DB.Model(&user).Updates(User{Phone: user.Phone, Email: user.Email}).Error; err != nil {
		err := fmt.Errorf("更新失败了")
		fmt.Println(err.Error())
		return false
	}
	return true
}
