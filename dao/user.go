package dao

import "log"

type User struct {
	ID         int64
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
}

func (User) TableName() string {
	return "users"
}

func init() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		return
	}
}

func SaveUser(user *User) {
	err := DB.Create(user).Error
	if err != nil {
		log.Panicln("insert user error:", err)
	}
}

func GetUserById(id int64) User {
	var user User
	err := DB.Where("id=?", id).First(&user).Error
	if err != nil {
		log.Panicln("get user by id error:", err)
	}
	return user
}

func GetAllUser() []User {
	var users []User
	err := DB.Find(&users).Error
	if err != nil {
		log.Panicln("get users error:", err)
	}
	return users
}

func UpdateUser(id int64) {
	var user User
	err := DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Panicln("update user error:", err)
	}
	user.Username = "张三"
	DB.Save(&user)
}
