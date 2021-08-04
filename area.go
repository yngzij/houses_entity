package houses_entity

import "time"

type Area struct {
	Id        int        `gorm:"id"`
	Name      string     `gorm:"Name"`
	Code      string     `gorm:"code"`
	Tag       int        `gorm:"tag"`
	DeletedAt *time.Time `gorm:"deleted_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	CreatedAt time.Time  `gorm:"created_at"`
	AdvertId  int        `gorm:"advert_id"`
}

func (Area) TableName() string {
	return "areas"
}
