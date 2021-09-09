package houses_entity

import "time"

type User struct {
	Id          int        `gorm:"id"`
	UserID      string     `gorm:"user_id"`
	AdvertId    int        `gorm:"advert_id"`
	Avatar      string     `gorm:"avatar"`
	NickName    string     `gorm:"nick_name"`
	Phone       string     `gorm:"phone"`
	Gender      int        `gorm:"gender"`
	Country     string     `gorm:"country"`
	City        string     `gorm:"city"`
	ChatGroupId string     `gorm:"chat_group_id"`
	UserId      string     `gorm:"user_id"`
	OpenId      string     `gorm:"open_id"`
	RealName    string     `gorm:"real_name"`
	UpdatedAt   time.Time  `gorm:"updated_at"`
	CreatedAt   time.Time  `gorm:"created_at"`
	DeletedAt   *time.Time `gorm:"deleted_at"`
}

func (u User) TableName() string {
	return "users"
}
