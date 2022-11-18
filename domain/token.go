package domain

type Token struct {
	Id          int    `gorm:"PrimaryKey"`
	Email       string `db:"email" json:"email"`
	Token       string `db:"token" json:"token"`
	DateCreated int64  `db:"date_created"`
}

func (Token) TableName() string {
	return "user_token"
}
