package model

// User of application
type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
}

func (user *User) IsValid() bool {
	return len(user.ID) > 0 && len(user.FirstName) > 0 && len(user.LastName) > 0
}
