package houses_entity

import "time"

type Advert struct {
	Id        int        `gorm:"id"`
	UsedNum   int        `gorm:"used_num"`
	UnusedNum int        `gorm:"unused_num"`
	UserId    int        `gorm:"user_id"`
	TotalNum  int        `gorm:"total_num"`
	Status    int        `gorm:"status"`
	Sort      int        `gorm:"sort"`
	DeletedAt *time.Time `gorm:"deleted_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	CreatedAt time.Time  `gorm:"created_at"`
}

func (a Advert) TableName() string {
	return "adverts"
}
