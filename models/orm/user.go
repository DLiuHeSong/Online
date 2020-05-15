package orm

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	NickName   string
	Cellphone  string
	Sex        int64
	WeChatName string
	PWD        string
	AvatarURL  string
}

func (u *User) TableName() string {
	return "user"
}
