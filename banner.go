package houses_entity

type Banner struct {
	Id       int    `gorm:"id"`
	ImageURL string `gorm:"image_url"`
}
