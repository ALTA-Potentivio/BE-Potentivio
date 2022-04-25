package auth

type AuthRepositoryInterface interface {
	LoginArtist(email string, password string) (string, uint, string, error)
	LoginCafe(email string, password string) (string, uint, string, error)
}
