package model

import "time"

// User represents schema of user record in a database
type User struct {
	ID           int    `gorm:"column:id"`
	CreatedAt    *int64 `gorm:"column:created_at"`
	UpdatedAt    *int64 `gorm:"column:updated_at"`
	DeletedAt    *int64 `gorm:"column:deleted_at"`
	EmailAddress string `gorm:"column:email_address"`
	Password     string `gorm:"column:password"`
	Salt         string `gorm:"column:salt"`
	FirstName    string `gorm:"column:first_name"`
	LastName     string `gorm:"column:last_name"`
	IsActive     int    `gorm:"column:is_active"`
}

type UserAccount struct {
	ID        string     `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	Email     string     `gorm:"column:email;unique"`
	Password  string     `gorm:"column:password"`
	Salt      string     `gorm:"column:salt"`
	FirstName string     `gorm:"column:first_name"`
	LastName  string     `gorm:"column:last_name"`
	IsActive  bool       `gorm:"column:is_active"`
	OldID     int64      `gorm:"column:old_id"`
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "public.user"
}

func (UserAccount) TableName() string {
	return "public.user_accounts"
}
