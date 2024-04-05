package entity

type Distance struct {
	ID       string `gorm:"primary_key" json:"id"`
	UserID   string `gorm:"type:varchar(36);index" json:"user_id"`
	Distance int    `gorm:"distance" json:"distance"`
}
