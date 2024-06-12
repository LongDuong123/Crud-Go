package domain

type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Age      int    `json:"Age"`
	Gender   string `json:"Gender"`
	Email    string `json:"Email"`
	Role     string `json:"Role"`
	Password string `json:"Password"`
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	GetByPage(start int, end int) ([]User, error)
	GetByEmail(email string) (*User, error)
	CheckEmail(email string) (bool, error)
	Create(user *User) error
	UpdateByID(id int, user *User) error
	DeleteByID(id int) error
}
