package houses_entity

type Area struct {
	Id        int        `gorm:"id"`
	Value     string     `gorm:"value"`
	DeletedAt *time.Time `gorm:"deleted_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	CreatedAt time.Time  `gorm:"created_at"`
	AdvertId int    `gorm:"advert_id"`
}

func (Area)TableName() string {
	return "areas"
}
