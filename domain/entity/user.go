package entity

// User .
type User struct {
	Model
	Name string `json:"name" gorm:"index:idx_user_1"`
}
