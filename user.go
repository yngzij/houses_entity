package houses_entity

import "time"

type User struct {
	ID          int        `gorm:"id"`
	UserID      string     `gorm:"user_id"`
	AdvertID    int        `gorm:"advert_id"`
	Avatar      string     `gorm:"avatar"`
	NickName    string     `gorm:"nick_name"`
	Phone       string     `gorm:"phone"`
	Gender      int        `gorm:"gender"`
	Country     string     `gorm:"country"`
	City        string     `gorm:"city"`
	ChatGroupID string     `gorm:"chat_group_id"`
	OpenID      string     `gorm:"open_id;index:idx_uid"`
	RealName    string     `gorm:"real_name"`
	UpdatedAt   time.Time  `gorm:"updated_at"`
	CreatedAt   time.Time  `gorm:"created_at"`
	DeletedAt   *time.Time `gorm:"deleted_at"`
}

func (u User) TableName() string {
	return "users"
}
