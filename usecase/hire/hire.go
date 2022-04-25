package hire

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"os"
	"potentivio-app/entities"
	"potentivio-app/repository/artist"
	"potentivio-app/repository/cafe"
	"potentivio-app/repository/hire"
	"strings"
	"time"
)

type HireUseCase struct {
	HireRepository   hire.HireRepositoryInterface
	ArtistRepository artist.ArtistRepositoryInterface
	CafeRepository   cafe.CafeRepositoryInterface
}

func NewHireUseCase(hireRepo hire.HireRepositoryInterface, artistRepo artist.ArtistRepositoryInterface, cafeRepo cafe.CafeRepositoryInterface) HireUseCaseInterface {
	return &HireUseCase{
		HireRepository:   hireRepo,
		ArtistRepository: artistRepo,
		CafeRepository:   cafeRepo,
	}
}

func (huc *HireUseCase) CreateHire(hire entities.Hire) error {

	artistData, err := huc.ArtistRepository.GetArtistByIdForHire(hire.IdArtist)

	if err != nil {
		return errors.New("artist not found")

	}

	cafeData, _, _ := huc.CafeRepository.GetCafeById(int(hire.IdCafe))
	err = huc.HireRepository.CheckHire(hire)

	if err != nil {
		return errors.New("Artis not Available")

	}

	hire.Price = *artistData.Price
	hire.AccountNumberArtist = artistData.AccountNumber
	hire.AccountNumberCafe = cafeData.AccountNumber

	err = huc.HireRepository.CreateHire(hire)

	return err

}

func (huc *HireUseCase) GetHireByIdArtist(IdArtist int) ([]entities.Hire, error) {
	hires, err := huc.HireRepository.GetHireByIdArtist(IdArtist)
	return hires, err
}

func (huc *HireUseCase) GetHireByIdCafe(IdCafe int) ([]entities.Hire, error) {
	hires, err := huc.HireRepository.GetHireByIdCafe(IdCafe)
	return hires, err
}

func (huc *HireUseCase) AcceptHire(hire entities.Hire) error {

	hires, err := huc.HireRepository.GetHireById(int(hire.ID))
	if hires.StatusArtist == "waiting payment" || hires.IdArtist != hire.IdArtist {
		return errors.New("Failed to accept")
	}

	year, month, day := time.Now().Date()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()

	invoiceNum := fmt.Sprint("invoice/", hire.IdCafe, "/", hire.IdArtist, "/", year, month, day, hour, minute, second)
	invoiceNum = strings.ReplaceAll(invoiceNum, " ", "")
	hires.Invoice = invoiceNum

	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY_XENDIT")

	cafeData, _, err := huc.CafeRepository.GetCafeById(int(hires.IdCafe))

	if err != nil {
		return errors.New("Artis not Available")
	}

	customer := xendit.InvoiceCustomer{
		GivenNames: cafeData.Name,
		Email:      cafeData.Email,
	}

	NotificationType := []string{"email"}

	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
		InvoiceCreated: NotificationType,
		InvoicePaid:    NotificationType,
		InvoiceExpired: NotificationType,
	}

	data := invoice.CreateParams{
		ExternalID:                     hires.Invoice,
		Amount:                         hires.Price,
		Description:                    "Invoice Demo #123",
		InvoiceDuration:                300,
		Customer:                       customer,
		CustomerNotificationPreference: customerNotificationPreference,
		Currency:                       "IDR",
	}

	resp, err := invoice.Create(&data)

	if err != nil {
		log.Info(err)
	}

	var paymentUrl = resp.InvoiceURL
	hires.PaymentUrl = &paymentUrl
	hires.StatusArtist = "waiting payment"
	hires.StatusCafe = "waiting payment"

	err = huc.HireRepository.UpdateHire(int(hire.ID), hires)

	return err
}

func (huc *HireUseCase) CancelHireByCafe(hire entities.Hire) error {
	hires, err := huc.HireRepository.GetHireById(int(hire.ID))

	if hires.StatusArtist != "waiting" || hires.IdCafe != hire.IdCafe {
		return errors.New("Failed to cancel")
	}
	var id = int(hire.ID)

	hires.StatusArtist = "Canceled"
	hires.StatusCafe = "Canceled"
	hires.Comment = hire.Comment
	err = huc.HireRepository.UpdateHire(id, hires)
	return err
}
