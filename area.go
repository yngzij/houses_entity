package houses_entity

import (
	"encoding/json"
)

type JSON json.RawMessage

type Area struct {
	ID       int    `gorm:"id"`
	Name     string `gorm:"name" json:"name"`
	Code     string `gorm:"code" json:"code"`
	AdvertId int    `gorm:"advert_id"`
}

func (Area) TableName() string {
	return "areas"
}
