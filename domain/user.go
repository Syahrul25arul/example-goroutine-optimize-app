package domain

import "database/sql"

type User struct {
	Id          int64  `gorm:"PrimaryKey"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	Role        string `db:"role"`
	Email       string `db:"email"`
	EmailVerify string `db:"email_verify"`
	CreatedAt   int64  `db:"created_at"`
}

type UserRegisterRequest struct {
	Id       sql.NullInt64 `json:"id"`
	Username string        `json:"username" required:"true"`
	Password string        `json:"password" required:"true" min:"5"`
	Role     string        `json:"role" required:"true"`
	Email    string        `json:"email" required:"true" min:"5"`
}

type UserVerifyEmailRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (u *UserRegisterRequest) ToUser() User {
	user := User{
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
		Email:    u.Email,
	}

	if u.Id.Valid {
		user.Id = u.Id.Int64
	}

	return user
}
