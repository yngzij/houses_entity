package houses_entity

import "time"

type Node struct {
	Id       int       `gorm:"id"`
	ImageURL string    `gorm:"image_url"`
	Title    string    `gorm:"title"`
	Sort     int       `gorm:"sort"`
	Deleted  time.Time `gorm:"deleted"`
	Updated  time.Time `gorm:"updated"`
	Created  time.Time `gorm:"created"`
}

func (n Node) TableName() string {
	return "nodes"
}
