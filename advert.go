package houses_entity

import (
	"time"

	"github.com/lib/pq"
)

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
	ID          int            `gorm:"id"`
	UsedNum     int            `gorm:"used_num"`
	Avatar      string         `gorm:"avatar"`
	NickName    string         `gorm:"nick_name"`
	Area        pq.StringArray `gorm:"type:text[]"`
	Title       string         `gorm:"title"`
	Local       string         `gorm:"local"`
	ChatGroupID string         `gorm:"chat_group_id"`
	UserID      string         `gorm:"user_id"`
	Permission  int            `gorm:"permission"`
	OpenID      string         `gorm:"open_id"`
	TotalNum    int            `gorm:"total_num"`
	Desc        string         `gorm:"desc"`
	Status      int            `gorm:"status"`
	Sort        int            `gorm:"sort"`
	DeletedAt   *time.Time     `gorm:"deleted_at"`
	UpdatedAt   time.Time      `gorm:"updated_at"`
	CreatedAt   time.Time      `gorm:"created_at"`
	Phone       string         `gorm:"phone"`
}

func (a Advert) TableName() string {
	return "adverts"
}
