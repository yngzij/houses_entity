package houses_entity

import "time"

type Banner struct {
	Id        int       `gorm:"id"`
	ImageURL  string    `gorm:"image_url"`
	Sort      int       `gorm:"sort"`
	DeletedAt time.Time `gorm:"deleted_at"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func (b Banner) TableName() string {
	return "banners"
}
