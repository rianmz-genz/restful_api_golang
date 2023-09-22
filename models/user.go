package models

// Model User
type User struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    Username string `gorm:"unique" json:"username"`
    Password string `json:"-"`
}