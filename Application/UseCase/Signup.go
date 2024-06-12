package usecase

import (
	domain "crud/Domain"

	"golang.org/x/crypto/bcrypt"
)

type SignUpUseCase struct {
	Signup domain.UserRepository
}

func NewSignUpUseCase(sui domain.UserRepository) domain.SignupInteractor {
	return &SignUpUseCase{Signup: sui}
}

func (signUpUseCase *SignUpUseCase) Create(sgu *domain.Signup) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(sgu.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	sgu.Password = string(hashPassword)
	return signUpUseCase.Signup.Create(&domain.User{Name: sgu.Name, Email: sgu.Email, Password: sgu.Password})
}

func (signUpUseCase *SignUpUseCase) CheckUserByEmail(email string) (bool, error) {
	return signUpUseCase.Signup.CheckEmail(email)
}
