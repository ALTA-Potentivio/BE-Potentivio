package hire

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"potentivio-app/entities"
	"testing"
	"time"
)

func TestCreateHire(t *testing.T) {
	t.Run(
		"Create Succes", func(t *testing.T) {
			var mockHireRepository = MockupHireRepository{}
			var mockArtistRepository = MockupArtisRepository{}
			var mockCafeRepository = MockupCafeRepository{}

			hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)

			accountNumber := "123"
			var price float64 = 10000
			mockArtistRepository.On("GetArtistByIdForHire", mock.Anything).Return(entities.Artist{AccountNumber: &accountNumber, Price: &price}, nil)
			mockCafeRepository.On("GetCafeById", mock.Anything).Return(entities.Cafe{AccountNumber: &accountNumber}, 1, nil)
			mockHireRepository.On("CheckHire", mock.Anything).Return(entities.Hire{})
			mockHireRepository.On("CreateHire", mock.Anything).Return(nil)

			err := hireUseCase.CreateHire(entities.Hire{
				IdArtist:            1,
				IdCafe:              1,
				Invoice:             "123",
				Date:                time.Now(),
				StatusArtist:        "waiting",
				StatusCafe:          "waiting",
				Price:               price,
				AccountNumberArtist: &accountNumber,
				AccountNumberCafe:   &accountNumber,
			})
			assert.Nil(t, err)
		})
	t.Run(
		"Fail artist not found", func(t *testing.T) {
			var mockHireRepository = MockupHireRepository{}
			var mockArtistRepository = MockupArtisRepository{}
			var mockCafeRepository = MockupCafeRepository{}

			hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)

			mockArtistRepository.On("GetArtistByIdForHire", mock.Anything).Return(entities.Artist{}, errors.New("artist not found"))

			err := hireUseCase.CreateHire(entities.Hire{})
			assert.Error(t, err)
		})
	t.Run(
		"Fail artist not available", func(t *testing.T) {
			var mockHireRepository = MockupHireRepository{}
			var mockArtistRepository = MockupArtisRepository{}
			var mockCafeRepository = MockupCafeRepository{}

			hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)

			mockArtistRepository.On("GetArtistByIdForHire", mock.Anything).Return(entities.Artist{}, nil)
			mockCafeRepository.On("GetCafeById", mock.Anything).Return(entities.Cafe{}, 1, nil)
			mockHireRepository.On("CheckHire", mock.Anything).Return(entities.Hire{StatusArtist: "waiting"})

			err := hireUseCase.CreateHire(entities.Hire{})
			assert.Error(t, err)
		})

	t.Run(
		"Fail, because artist data not complete", func(t *testing.T) {
			var mockHireRepository = MockupHireRepository{}
			var mockArtistRepository = MockupArtisRepository{}
			var mockCafeRepository = MockupCafeRepository{}

			hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)

			mockArtistRepository.On("GetArtistByIdForHire", mock.Anything).Return(entities.Artist{}, nil)
			mockCafeRepository.On("GetCafeById", mock.Anything).Return(entities.Cafe{}, 1, nil)
			mockHireRepository.On("CheckHire", mock.Anything).Return(entities.Hire{})

			err := hireUseCase.CreateHire(entities.Hire{})
			assert.Error(t, err)
		})

	t.Run(
		"Fail, because cafe data not complete", func(t *testing.T) {
			var mockHireRepository = MockupHireRepository{}
			var mockArtistRepository = MockupArtisRepository{}
			var mockCafeRepository = MockupCafeRepository{}

			hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)
			var accountNumber = "123"
			var price float64 = 10000

			mockArtistRepository.On("GetArtistByIdForHire", mock.Anything).Return(entities.Artist{AccountNumber: &accountNumber, Price: &price}, nil)
			mockCafeRepository.On("GetCafeById", mock.Anything).Return(entities.Cafe{}, 1, nil)
			mockHireRepository.On("CheckHire", mock.Anything).Return(entities.Hire{})

			err := hireUseCase.CreateHire(entities.Hire{})
			assert.Error(t, err)
		})

}

func TestGetHireByIdArtist(t *testing.T) {
	t.Run("Get hire by id artis succes", func(t *testing.T) {
		var mockHireRepository = MockupHireRepository{}
		var mockArtistRepository = MockupArtisRepository{}
		var mockCafeRepository = MockupCafeRepository{}

		hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)
		mockHireRepository.On("GetHireByIdArtist", mock.Anything).Return([]entities.Hire{{IdArtist: 1}}, nil)

		hire, err := hireUseCase.GetHireByIdArtist(1)
		assert.Equal(t, uint(1), hire[0].IdArtist)
		assert.Nil(t, err)
	})

}

func TestGetHireByIdCafe(t *testing.T) {
	t.Run(
		"Get hire by id cafe succes", func(t *testing.T) {
			var mockHireRepository = MockupHireRepository{}
			var mockArtistRepository = MockupArtisRepository{}
			var mockCafeRepository = MockupCafeRepository{}

			hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)
			mockHireRepository.On("GetHireByIdCafe", mock.Anything).Return([]entities.Hire{{IdCafe: 1}}, nil)

			hire, err := hireUseCase.GetHireByIdCafe(1)
			assert.Equal(t, uint(1), hire[0].IdCafe)
			assert.Nil(t, err)
		})

}

func TestRejectHire(t *testing.T) {
	t.Run(
		"Reject hire succes", func(t *testing.T) {
			var mockHireRepository = MockupHireRepository{}
			var mockArtistRepository = MockupArtisRepository{}
			var mockCafeRepository = MockupCafeRepository{}

			hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)
			mockHireRepository.On("GetHireById", mock.Anything).Return(entities.Hire{StatusArtist: "waiting"}, nil)
			mockHireRepository.On("UpdateHire", mock.Anything, mock.Anything).Return(nil)

			err := hireUseCase.Rejecthire(entities.Hire{})
			assert.Nil(t, err)
		})

	t.Run(
		"Reject hire failed", func(t *testing.T) {
			var mockHireRepository = MockupHireRepository{}
			var mockArtistRepository = MockupArtisRepository{}
			var mockCafeRepository = MockupCafeRepository{}

			hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)
			mockHireRepository.On("GetHireById", mock.Anything).Return(entities.Hire{}, nil)
			mockHireRepository.On("UpdateHire", mock.Anything, mock.Anything).Return(nil)

			err := hireUseCase.Rejecthire(entities.Hire{})
			assert.NotNil(t, err)
		})

}

func TestRating(t *testing.T) {
	t.Run(
		"Rating succes", func(t *testing.T) {
			var mockHireRepository = MockupHireRepository{}
			var mockArtistRepository = MockupArtisRepository{}
			var mockCafeRepository = MockupCafeRepository{}

			hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)
			mockHireRepository.On("GetHireById", mock.Anything).Return(entities.Hire{StatusArtist: "done", StatusCafe: "done"}, nil)
			mockHireRepository.On("UpdateHire", mock.Anything, mock.Anything).Return(nil)
			mockHireRepository.On("Rating", mock.Anything).Return(nil)

			err := hireUseCase.Rating(entities.Hire{})
			assert.Nil(t, err)
		})

	t.Run(
		"Rating hire failed", func(t *testing.T) {
			var mockHireRepository = MockupHireRepository{}
			var mockArtistRepository = MockupArtisRepository{}
			var mockCafeRepository = MockupCafeRepository{}

			hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)
			mockHireRepository.On("GetHireById", mock.Anything).Return(entities.Hire{}, nil)
			mockHireRepository.On("UpdateHire", mock.Anything, mock.Anything).Return(nil)
			mockHireRepository.On("Rating", mock.Anything).Return(nil)

			err := hireUseCase.Rating(entities.Hire{})
			assert.NotNil(t, err)
		})

	t.Run(
		"Rating failed because you have given a rating ", func(t *testing.T) {
			var mockHireRepository = MockupHireRepository{}
			var mockArtistRepository = MockupArtisRepository{}
			var mockCafeRepository = MockupCafeRepository{}

			hireUseCase := NewHireUseCase(&mockHireRepository, &mockArtistRepository, &mockCafeRepository)
			mockHireRepository.On("GetHireById", mock.Anything).Return(entities.Hire{StatusArtist: "done", StatusCafe: "done", Rating: 4}, nil)
			mockHireRepository.On("UpdateHire", mock.Anything, mock.Anything).Return(nil)
			mockHireRepository.On("Rating", mock.Anything).Return(nil)

			err := hireUseCase.Rating(entities.Hire{})
			assert.NotNil(t, err)
		})

}
