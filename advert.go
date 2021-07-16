package houses_entity

import "time"

const (
	AdvertAuditSuccess = 1 // 审核成功
	AdvertAuditFailure = 2 // 审核失败
	AdvertAuditing     = 3 // 审核中
)

type Advert struct {
	Id          int        `gorm:"id"`
	UsedNum     int        `gorm:"used_num"`
	Avatar      string     `gorm:"avatar"`
	NickName    string     `gorm:"nick_name"`
	Title       string     `gorm:"title"`
	UnusedNum   int        `gorm:"unused_num"`
	ChatGroupId string     `gorm:"chat_group_id"`
	UserId      int        `gorm:"user_id"`
	OpenId      string     `gorm:"open_id"`
	TotalNum    int        `gorm:"total_num"`
	Desc        string     `gorm:"desc"`
	Status      int        `gorm:"status"`
	Sort        int        `gorm:"sort"`
	DeletedAt   *time.Time `gorm:"deleted_at"`
	UpdatedAt   time.Time  `gorm:"updated_at"`
	CreatedAt   time.Time  `gorm:"created_at"`
	Phone       string     `gorm:"phone"`
	StartTime   int64      `gorm:"start_time"`
	EndTime     int64      `gorm:"end_time"`
}

func (a Advert) TableName() string {
	return "adverts"
}
