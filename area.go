package houses_entity

type Area struct {
	Id       int    `gorm:"id"`
	Value    string `gorm:"value"`
	AdvertId int    `gorm:"advert_id"`
}

func (Area)TableName() string {
	return "areas"
}
