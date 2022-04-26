package videoArtist

import _entities "potentivio-app/entities"

type VideoRepositoryInterface interface {
	PostVideo(video _entities.VideoArtist) error
	DeleteVideo(id int, idToken int) error
}
