package houses_entity

import (
	"time"

	"gorm.io/gorm"

	"github.com/lib/pq"
)

type AdvertStatus int

const (
	AdvertExpired    AdvertStatus = 1
	AdvertPublishing AdvertStatus = 2
	AdvertFinished   AdvertStatus = 3
	AdvertDeleted    AdvertStatus = 4
)

type Advert struct {
	ID          int            `gorm:"id, uniqueIndex"`
	Avatar      string         `gorm:"avatar"`
	NickName    string         `gorm:"nick_name"`
	Area        pq.StringArray `gorm:"type:text[]"`
	Local       string         `gorm:"local"`
	ChatGroupID string         `gorm:"chat_group_id"`
	UserID      string         `gorm:"user_id;index:idx_uid,type:btree"`
	Permissions int            `gorm:"permissions"`
	OpenID      string         `gorm:"open_id;index:idx_oid,type:btree"`
	Desc        string         `gorm:"desc"`
	Status      AdvertStatus   `gorm:"status;default:2;index:idx_satuts,type:btree"`
	Sort        int            `gorm:"column:sort;type:bigserial"`
	DeletedAt   gorm.DeletedAt `gorm:"deleted_at"`
	UpdatedAt   time.Time      `gorm:"updated_at"`
	CreatedAt   time.Time      `gorm:"created_at"`
	ExpiredAt   time.Time      `gorm:"expired_at"`
	Phone       string         `gorm:"phone"`
}

func (a Advert) TableName() string {
	return "adverts"
}
