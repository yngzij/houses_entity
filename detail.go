package houses_entity

import "time"

type Detail struct {
	Id       int       `gorm:"id"`
	ImageURL string    `gorm:"image_url"`
	HouseId  int       `gorm:"house_id"`
	Deleted  time.Time `gorm:"deleted"`
	Sort     int       `gorm:"sort"`
	Created  time.Time `gorm:"created"`
	Updated  time.Time `gorm:"updated"`
}

func (d Detail) TableName() string {
	return "details"
}
