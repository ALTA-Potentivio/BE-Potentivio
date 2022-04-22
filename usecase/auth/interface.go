package auth

type AuthUseCaseInterface interface {
	LoginArtist(email string, password string) (string, uint, error)
	LoginCafe(email string, password string) (string, uint, error)
}
