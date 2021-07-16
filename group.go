package houses_entity

import "time"

type Group struct {
	Id          int        `gorm:"id"`
	Phone       string     `gorm:"phone"`
	ChatGroupId string     `gorm:"chat_group_id"`
	AdvertId    int        `gorm:"advert_id"`
	OpenId      string     `gorm:"open_id"`
	DeletedAt   *time.Time `gorm:"deleted_at"`
	UpdatedAt   time.Time  `gorm:"updated_at"`
	CreatedAt   time.Time  `gorm:"created_at"`
}

func (g Group) TableName() string {
	return "groups"
}
