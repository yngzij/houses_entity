package houses_entity

import "time"

type Detail struct {
	ID        int        `gorm:"id"`
	ImageURL  string     `gorm:"image_url"`
	HouseID   int        `gorm:"house_id"`
	Sort      int        `gorm:"sort"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}

func (d Detail) TableName() string {
	return "details"
}
