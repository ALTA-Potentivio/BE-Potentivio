package videoArtist

import _entities "potentivio-app/entities"

type VideoUseCaseInterface interface {
	PostVideo(video _entities.VideoArtist, idToken int) error
	DeleteVideo(id int, idToken int) error
}
