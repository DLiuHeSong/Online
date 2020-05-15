package orm

import "github.com/jinzhu/gorm"

type UserInfo struct {
	gorm.Model
	UserID        int64
	LabelIDS      int64
	FocusLabelIDS string
	Medal         int64
	MediaUrlID    int64
	JoinEventIDS  string
}

func (ui *UserInfo) TableName() string {
	return "user_info"
}
