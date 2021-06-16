package houses_entity

import "time"

type Banner struct {
	Id       int       `gorm:"id"`
	ImageURL string    `gorm:"image_url"`
	Sort     int       `gorm:"sort"`
	Deleted  time.Time `gorm:"deleted"`
	Created  time.Time `gorm:"created"`
	Updated  time.Time `gorm:"updated"`
}

func (b Banner) TableName() string {
	return "banners"
}
