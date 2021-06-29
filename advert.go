package houses_entity

import "time"

type Advert struct {
	Id        int       `json:"id"`
	UsedNum   int       `json:"used_num"`
	UnusedNum int       `json:"unused_num"`
	TotalNum  int       `json:"total_num"`
	Status    int       `json:"status"`
	Sort      int       `json:"sort"`
	DeletedAt time.Time `gorm:"deleted_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	CreatedAt time.Time `gorm:"created_at"`
}
