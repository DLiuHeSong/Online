package orm

import "github.com/jinzhu/gorm"

type Label struct {
	gorm.Model
	LabelType int64
	Title     string
}

func (l *Label) TableName() string {
	return "label"
}
