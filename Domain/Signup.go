package domain

type Signup struct {
	Name     string
	Email    string
	Password string
}

type SignupInteractor interface {
	Create(signUp *Signup) error
	CheckUserByEmail(email string) (bool, error)
}
