package houses_entity

import "time"

type House struct {
	Id       int       `gorm:"id"`
	Price    int       `gorm:"price"`
	ImageUrl string    `gorm:"image_url"`
	NodeId   int       `gorm:"node_id"`
	Title    string    `gorm:"title"`
	Sort     int       `gorm:"sort"`
	Desc     string    `gorm:"desc"`
	Deleted  time.Time `gorm:"deleted"`
	Updated  time.Time `gorm:"updated"`
	Created  time.Time `gorm:"created"`
}

func (h House) TableName() string {
	return "houses"
}
