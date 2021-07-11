package houses_entity

type Group struct {
	Id          int    `gorm:"id"`
	Phone       string `gorm:"phone"`
	ChatGroupId string `gorm:"chat_group_id"`
}
