package auth

import (
	_authRepository "potentivio-app/repository/auth"
)

type AuthUseCase struct {
	authRepository _authRepository.AuthRepositoryInterface
}

func NewAuthUseCase(authRepo _authRepository.AuthRepositoryInterface) AuthUseCaseInterface {
	return &AuthUseCase{
		authRepository: authRepo,
	}
}

func (auc *AuthUseCase) LoginArtist(email string, password string) (string, uint, string, error) {
	token, idArtist, validation, err := auc.authRepository.LoginArtist(email, password)
	return token, idArtist, validation, err
}

func (auc *AuthUseCase) LoginCafe(email string, password string) (string, uint, string, error) {
	token, idCafe, validation, err := auc.authRepository.LoginCafe(email, password)
	return token, idCafe, validation, err
}
