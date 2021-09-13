package houses_entity

import "time"

const (
	AdvertAuditSuccess = 1 // 审核成功
	AdvertAuditFailure = 2 // 审核失败
	AdvertAuditing     = 3 // 审核中

	AdvertFnished  = 4 // 完成
	Adverting      = 5 // 发布中
	AdvertCanceled = 6 // 取消
	AdvertDeleted  = 7 // 已经删除
)

type Advert struct {
	ID      int `gorm:"id"`
	UsedNum int `gorm:"used_num"`
	Avatar      string     `gorm:"avatar"`
	NickName    string     `gorm:"nick_name"`
	Title       string     `gorm:"title"`
	AreaCode    string     `gorm:"area_code"`
	UnusedNum   int    `gorm:"unused_num"`
	ChatGroupID string `gorm:"chat_group_id"`
	UserId      int    `gorm:"user_id"`
	OpenID   string `gorm:"open_id"`
	TotalNum int    `gorm:"total_num"`
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
