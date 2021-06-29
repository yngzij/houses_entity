package houses_entity

import "time"

type House struct {
	Id        int       `gorm:"id"`
	Price     int       `gorm:"price"`
	ImageUrl  string    `gorm:"image_url"`
	NodeId    int       `gorm:"node_id"`
	Title     string    `gorm:"title"`
	Sort      int       `gorm:"sort"`
	Desc      string    `gorm:"desc"`
	DeletedAt time.Time `gorm:"deleted_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	CreatedAt time.Time `gorm:"created_at"`
}

func (h House) TableName() string {
	return "houses"
}
