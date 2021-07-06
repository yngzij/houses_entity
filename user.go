package houses_entity

import "time"

type User struct {
	Id         int        `gorm:"id"`
	AdvertId   int        `gorm:"advert_id"`
	Avatar     string     `gorm:"avatar"`
	NickName   string     `gorm:"nick_name"`
	Gender     int        `gorm:"gender"`
	Country    string     `gorm:"country"`
	City       string     `gorm:"city"`
	ChatroomId string     `gorm:"chatroom_id"`
	UserId     string     `gorm:"user_id"`
	OpenId     string     `gorm:"open_id"`
	RealName   string     `gorm:"real_name"`
	UpdatedAt  time.Time  `gorm:"updated_at"`
	CreatedAt  time.Time  `gorm:"created_at"`
	DeletedAt  *time.Time `gorm:"deleted_at"`
}

func (u User) TableName() string {
	return "users"
}
