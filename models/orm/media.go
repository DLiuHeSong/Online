package orm

import "github.com/jinzhu/gorm"

type Media struct {
	gorm.Model
	MediaURL string
}

func (m *Media) TableName() string {
	return "media"
}
