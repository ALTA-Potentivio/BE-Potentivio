package videoArtist

import (
	"errors"
	_entities "potentivio-app/entities"
	_videoRepository "potentivio-app/repository/videoArtist"
)

type VideoUseCase struct {
	videoRepository _videoRepository.VideoRepositoryInterface
}

func NewVideoUseCase(videoRepo _videoRepository.VideoRepositoryInterface) VideoUseCaseInterface {
	return &VideoUseCase{
		videoRepository: videoRepo,
	}
}

func (vuc *VideoUseCase) PostVideo(video _entities.VideoArtist, idToken int, name string) error {
	video.IdArtist = uint(idToken)
	if video.VideoUrl == "" {
		return errors.New("please insert link video")
	}
	error := vuc.videoRepository.PostVideo(video, name)
	return error
}

func (vuc *VideoUseCase) DeleteVideo(id int, idToken int, name string) error {
	error := vuc.videoRepository.DeleteVideo(id, idToken, name)
	return error
}
