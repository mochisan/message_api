package entity

// GroupUser .
type GroupUser struct {
	Model
	GroupID uint `json:"group_id" gorm:"index:idx_group_user_1; not null"`
	UserID  uint `json:"user_id" gorm:"index:idx_group_user_2; not null"`
}
