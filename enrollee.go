package houses_entity

import "time"

type Advert struct {
	Id        int       `gorm:"id"`
	UsedNum   int       `gorm:"use_num"`
	UnusedNum int       `gorm:"unused_num"`
	TotalNum  int       `gorm:"total_num"`
	Status    int       `gorm:"status"`
	Sort      int       `gorm:"sort"`
	Deleted   time.Time `gorm:"deleted"`
	Updated   time.Time `gorm:"updated"`
	Created   time.Time `gorm:"created"`
}

func (a Advert) TableName() string {
	return "adverts"
}
