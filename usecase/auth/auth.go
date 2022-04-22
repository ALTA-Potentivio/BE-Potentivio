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

func (auc *AuthUseCase) LoginArtist(email string, password string) (string, uint, error) {
	token, idArtist, err := auc.authRepository.LoginArtist(email, password)
	return token, idArtist, err
}

func (auc *AuthUseCase) LoginCafe(email string, password string) (string, uint, error) {
	token, idCafe, err := auc.authRepository.LoginCafe(email, password)
	return token, idCafe, err
}
