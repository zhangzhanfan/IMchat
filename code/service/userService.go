package service

import (
	"code/config"
	"code/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 获取用户列表
// @Produce  json
// @Tags 用户
// @Success 200 {string} json {"code", "data"}
// @Failure 500 {string} json {"code", "data"}
// @Router /user/GetUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	config.Success(data, c)
}

//创建用户
func CreateUser(c *gin.Context) {
	var userLogicInfo struct {
		Name       string `json:"name"`
		Password   string `json:"password"`
		Repassword string `json:"repassword"`
	}
	err := c.BindJSON(&userLogicInfo)
	if err != nil {
		return
	}
	if userLogicInfo.Password != userLogicInfo.Repassword {
		message := map[string]string{
			"message": "两次密码不一致1",
		}
		config.Failed(message, c)
		return
		// config.Failed("两次密码不一致", c)
	}
	userLogicInfo.Password = EncryptPassword(userLogicInfo.Password)
	user := models.User{
		Name:     userLogicInfo.Name,
		Password: userLogicInfo.Password,
	}
	models.CreateUser(user)
	data := map[string]interface{}{
		"message": "新增成功",
		"data":    user,
	}
	config.Success(data, c)
}

//删除用户
func DeleteUser(c *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	config.Success(id, c)
}

// 更新用户
func UpdateUser(c *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	fmt.Println(id)
	name := c.PostForm("name")
	fmt.Println(name)
	fmt.Println("wdwdwdwd我顶我顶我顶")
	password := c.PostForm("password")
	user.ID = uint(id)
	user.Name = name
	user.Password = password
	models.UpdateUser(user)
	config.Success("更新成功", c)
}

func UpdatePhoneAndEmail(c *gin.Context) {
	var body struct {
		Id    string `json:"id"`
		Phone string `json:"phone"`
		Email string `json:"email"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		return
	}
	phoneResult := CheckPhone(body.Phone)
	if !phoneResult {
		config.Failed("手机校验不通过", c)
		return
	}
	emailResult := CheckEmail(body.Email)
	if !emailResult {
		config.Failed("邮箱校验不通过", c)
		return
	}
	id, _ := strconv.Atoi(body.Id)
	var user models.User
	user.ID = uint(id)
	user.Phone = body.Phone
	user.Email = body.Email
	response := models.UpdateInfo(user)
	if !response {
		config.Failed("更新失败了", c)
		return
	}
	config.Success("更新成功！", c)
}
