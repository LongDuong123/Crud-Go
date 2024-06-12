package usecase

import domain "crud/Domain"

type ProfileUseCase struct {
	userRepository domain.UserRepository
}

func NewProfileUseCase(pu domain.UserRepository) domain.ProfileInteractor {
	return &ProfileUseCase{userRepository: pu}
}

func (pu *ProfileUseCase) GetProfileByID(id int) (*domain.Profile, error) {
	Profile, err := pu.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &domain.Profile{Name: Profile.Name, Age: Profile.Age, Gender: Profile.Gender, Email: Profile.Email}, nil
}

func (pu *ProfileUseCase) UpdateProfile(id int, profile *domain.Profile) error {
	profileUser := &domain.User{ID: id, Name: profile.Name, Age: profile.Age, Gender: profile.Gender, Email: profile.Email}
	return pu.userRepository.UpdateByID(id, profileUser)
}
