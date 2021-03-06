package hire

import (
	"errors"
	"fmt"
	"os"
	"potentivio-app/entities"
	"potentivio-app/repository/artist"
	"potentivio-app/repository/cafe"
	"potentivio-app/repository/hire"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursement"
	"github.com/xendit/xendit-go/invoice"
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
	hireData := huc.HireRepository.CheckHire(hire)

	if hireData.StatusArtist == "waiting" || hireData.StatusArtist == "waiting payment" || hireData.StatusArtist == "PAID" {
		return errors.New("Artis not Available")
	}

	if artistData.AccountNumber == nil || artistData.Price == nil {
		return errors.New("Can't hire artis, because artist data not complete")
	}

	if cafeData.AccountNumber == nil {
		return errors.New("Can't hire artis, because cafe data not complete")
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

	invoiceNum := fmt.Sprint("invoice/", hires.IdCafe, "/", hires.IdArtist, "/", year, month, day, hour, minute, second)
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
		InvoiceDuration:                3000,
		Customer:                       customer,
		CustomerNotificationPreference: customerNotificationPreference,
		Currency:                       "IDR",
	}

	resp, _ := invoice.Create(&data)

	//if err != nil {
	//	log.Info(err)
	//}

	var paymentUrl = resp.InvoiceURL
	hires.PaymentUrl = paymentUrl
	hires.StatusArtist = "waiting payment"
	hires.StatusCafe = "waiting payment"
	hires.IDXendit = resp.ID

	err = huc.HireRepository.UpdateHire(int(hire.ID), hires)

	return err
}

func (huc *HireUseCase) CancelHireByCafe(hire entities.Hire) error {
	hires, err := huc.HireRepository.GetHireById(int(hire.ID))

	if hires.StatusArtist == "PAID" || hires.StatusArtist == "EXPIRED" || hires.IdCafe != hire.IdCafe {
		return errors.New("Failed to cancel")
	}
	var id = int(hire.ID)

	if hires.Invoice != "" {
		xendit.Opt.SecretKey = os.Getenv("SECRET_KEY_XENDIT")

		data := invoice.ExpireParams{
			ID: hires.IDXendit,
		}

		_, err = invoice.Expire(&data)
		if err != nil {
			log.Info(err)
		}
	}

	hires.StatusArtist = "Canceled"
	hires.StatusCafe = "Canceled"
	hires.Comment = hire.Comment

	err = huc.HireRepository.UpdateHire(id, hires)
	if err != nil {
		return err
	}
	err = huc.HireRepository.DeleteHire(hire)
	return err
}

func (huc *HireUseCase) Rejecthire(hire entities.Hire) error {
	hires, err := huc.HireRepository.GetHireById(int(hire.ID))

	if hires.StatusArtist != "waiting" || hires.IdArtist != hire.IdArtist {
		return errors.New("Failed to Reject")
	}
	var id = int(hire.ID)

	hires.StatusArtist = "Rejected"
	hires.StatusCafe = "Rejected"

	err = huc.HireRepository.UpdateHire(id, hires)
	return err
}

func (huc *HireUseCase) CancelHireByArtis(hire entities.Hire) error {
	hires, err := huc.HireRepository.GetHireById(int(hire.ID))

	if hires.StatusArtist != "PAID" || hires.IdArtist != hire.IdArtist {
		return errors.New("Failed to cancel")

	}

	cafe, _ := huc.CafeRepository.GetCafeByIdForHire(hires.IdCafe)

	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY_XENDIT")

	year, month, day := time.Now().Date()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()

	disbursementKey := fmt.Sprint("disbursement/", hires.IdCafe, "/", hires.IdArtist, "/", year, month, day, hour, minute, second)
	disbursementKey = strings.ReplaceAll(disbursementKey, " ", "")

	createData := disbursement.CreateParams{
		IdempotencyKey:    disbursementKey,
		ExternalID:        hires.Invoice,
		BankCode:          "BCA",
		AccountHolderName: cafe.Name,
		AccountNumber:     *hires.AccountNumberCafe,
		Description:       "Pengembalian dana dari Potentivio",
		Amount:            hires.Price,
		EmailTo:           []string{cafe.Email},
	}

	_, err = disbursement.Create(&createData)
	if err != nil {
		log.Info(err)
	}

	hires.StatusArtist = "canceled"
	hires.StatusCafe = "canceled"

	err = huc.HireRepository.UpdateHire(int(hires.ID), hires)

	return err

}
func (huc *HireUseCase) Rating(hire entities.Hire) error {

	hires, err := huc.HireRepository.GetHireById(int(hire.ID))
	if hires.StatusArtist != "done" || hires.StatusCafe != "done" {
		return errors.New("status not done")
	}

	if hires.Rating != 0 {
		return errors.New("you have given a rating")
	}
	var id = int(hire.ID)
	hires.Comment = hire.Comment
	hires.Rating = hire.Rating
	err = huc.HireRepository.UpdateHire(id, hires)

	var rating entities.Rating

	rating.IdArtist = hires.IdArtist
	rating.Rating = hires.Rating

	err = huc.HireRepository.Rating(rating)
	return err

}

func (huc *HireUseCase) CallBack(hire entities.Hire) error {
	err := huc.HireRepository.CallBack(hire)
	return err

}

func (huc *HireUseCase) Done(hire entities.Hire) error {
	hires, err := huc.HireRepository.GetHireById(int(hire.ID))

	if hires.StatusArtist != "PAID" || hires.IdCafe != hire.IdCafe {
		return errors.New("Failed to done")

	}

	artist, _ := huc.ArtistRepository.GetArtistByIdForHire(hires.IdArtist)

	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY_XENDIT")

	year, month, day := time.Now().Date()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()

	disbursementKey := fmt.Sprint("disbursement/", hires.IdCafe, "/", hires.IdArtist, "/", year, month, day, hour, minute, second)
	disbursementKey = strings.ReplaceAll(disbursementKey, " ", "")

	createData := disbursement.CreateParams{
		IdempotencyKey:    disbursementKey,
		ExternalID:        hires.Invoice,
		BankCode:          "BCA",
		AccountHolderName: artist.Name,
		AccountNumber:     *hires.AccountNumberArtist,
		Description:       "Pembayaran dana dari Potentivio",
		Amount:            hires.Price,
		EmailTo:           []string{artist.Email},
	}

	_, err = disbursement.Create(&createData)
	if err != nil {
		log.Info(err)
	}

	hires.StatusArtist = "done"
	hires.StatusCafe = "done"

	err = huc.HireRepository.UpdateHire(int(hires.ID), hires)

	return err
}
