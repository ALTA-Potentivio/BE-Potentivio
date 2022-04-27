package videoArtist

import _entities "potentivio-app/entities"

type VideoRepositoryInterface interface {
	PostVideo(video _entities.VideoArtist, name string) error
	DeleteVideo(id int, idToken int, name string) error
}
