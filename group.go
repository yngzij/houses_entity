package houses_entity

import "time"

type Group struct {
	ID          int        `gorm:"id"`
	Phone       string     `gorm:"phone"`
	ChatGroupID string     `gorm:"chat_group_id"`
	AdvertID    int        `gorm:"advert_id"`
	UserID      string     `gorm:"user_id"`
	Avatar      string     `gorm:"avatar"`
	DeletedAt   *time.Time `gorm:"deleted_at"`
	UpdatedAt   time.Time  `gorm:"updated_at"`
	CreatedAt   time.Time  `gorm:"created_at"`
}

func (g Group) TableName() string {
	return "groups"
}
