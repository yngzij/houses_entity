package houses_entity

import "time"

type Area struct {
	Id        int        `gorm:"id"`
	Value     string     `gorm:"value"`
	DeletedAt *time.Time `gorm:"deleted_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	CreatedAt time.Time  `gorm:"created_at"`
}
