package domain

type Article struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      uint   `json:"user_id"`
	User        User   `json:"user" gorm:"foreignKey:UserId"`
}
