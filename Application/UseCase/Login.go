package usecase

import (
	domain "crud/Domain"
	"crud/internal"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	Repository domain.UserRepository
}

func NewLoginUseCase(rp domain.UserRepository) domain.LoginInteractor {
	return &LoginUseCase{Repository: rp}
}

func (lu *LoginUseCase) GetByEmail(email string) (*domain.User, error) {
	return lu.Repository.GetByEmail(email)
}
func (lu *LoginUseCase) CheckPassword(hashpassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password)) == nil
}

func (lu *LoginUseCase) CreateAccessToken(id int, userName string, expiry time.Time) (string, error) {
	return internal.CreateToken(id, userName, expiry)
}
func (lu *LoginUseCase) CreateRefreshToken(id int, userName string, expiry time.Time) (string, error) {
	return internal.CreateToken(id, userName, expiry)
}
