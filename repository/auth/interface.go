package auth

type AuthRepositoryInterface interface {
	LoginArtist(email string, password string) (string, uint, error)
	LoginCafe(email string, password string) (string, uint, error)
}
