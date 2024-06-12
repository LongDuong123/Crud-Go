package domain

type Profile struct {
	Name   string `json:"Name"`
	Age    int    `json:"Age"`
	Gender string `json:"Gender"`
	Email  string `json:"Email"`
}

type ProfileInteractor interface {
	GetProfileByID(id int) (*Profile, error)
	UpdateProfile(id int, user *Profile) error
}
