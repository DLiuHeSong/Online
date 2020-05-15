package orm

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	EventType    int64
	Title        int64
	Introduction string
	OrganizerID  int64
	ManagerIDS   string
	Longitude    float64
	Latitude     float64
	MinSize      int64
	MaxSize      int64
	Location     string
	LabelIDS     string
	StartTime    int64
	EndTime      int64
	State        int64
	UserIDS      int64
}

func (e *Event) TableName() string {
	return "event"
}
