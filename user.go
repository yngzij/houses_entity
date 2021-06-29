package houses_entity

import "time"

type User struct {
	Id        int       `gorm:"id"`
	AdvertId  int       `gorm:"advert_id"`
	Avatar    string    `gorm:"avatar"`
	NickName  string    `gorm:"nick_name"`
	RealName  string    `gorm:"real_name"`
	Local     string    `gorm:"local"`
	DeletedAt time.Time `gorm:"deleted_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	CreatedAt time.Time `gorm:"created_at"`
}

func (u User) TableName() string {
	return "users"
}
