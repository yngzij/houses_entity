package houses_entity

import "time"

type Node struct {
	Id        int        `gorm:"id"`
	ImageURL  string     `gorm:"image_url"`
	Title     string     `gorm:"title"`
	Sort      int        `gorm:"sort"`
	DeletedAt *time.Time `gorm:"deleted_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	CreatedAt time.Time  `gorm:"created_at"`
}

func (n Node) TableName() string {
	return "nodes"
}
