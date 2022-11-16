package domain

import "time"

type User struct {
	Username    string    `gorm:"PrimaryKey" json:"username" required:"true"`
	Password    string    `db:"password" json:"password" required:"true" min:"5"`
	Role        string    `db:"role" json:"role" required:"true"`
	Email       string    `db:"email" json:"email" required:"true" min:"5"`
	EmailVerify string    `db:"email_verify" json:"email_verify" required:"true"`
	CreatedAt   time.Time `db:"created_at"`
}
