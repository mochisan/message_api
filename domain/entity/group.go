package entity

// Group .
type Group struct {
	Model
	Name string `json:"name"`

	Users []User `json:"users"`
}
