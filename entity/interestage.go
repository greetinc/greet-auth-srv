package entity

type InterestAge struct {
	ID     string `gorm:"primary_key" json:"id"`
	UserID string `gorm:"type:varchar(36);index" json:"user_id"`
	// User   User   `json:"user" gorm:"foreignKey:UserID"`
	MinAge int `gorm:"min_age" json:"min_age"`
	MaxAge int `gorm:"max_age" json:"max_age"`
}
