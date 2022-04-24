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

	artistData, _, _, _, err := huc.ArtistRepository.GetArtistById(hire.IdArtist)
	if err != nil {
		return errors.New("Artist not found")

	}
	//var AccountNumber = "1234"
	//
	//var artistData = entities.Artist{
	//
	//	Price:         10000,
	//	AccountNumber: &AccountNumber,
	//}
	cafeData, _, _ := huc.CafeRepository.GetCafeById(int(hire.IdCafe))
	err = huc.HireRepository.CheckHire(hire)
	fmt.Println(err)
	if err == nil {
		return errors.New("Artis not Available")
	}

	hire.Price = artistData.Price
	hire.AccountNumberArtist = artistData.AccountNumber
	hire.AccountNumberCafe = cafeData.AccountNumber

	year, month, day := time.Now().Date()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()

	invoiceNum := fmt.Sprint("invoice/", hire.IdCafe, "/", hire.IdArtist, "/", year, month, day, hour, minute, second)

	hire.Invoice = invoiceNum
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY_XENDIT")

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
		ExternalID:                     hire.Invoice,
		Amount:                         hire.Price,
		Description:                    "Invoice Demo #123",
		InvoiceDuration:                300,
		Customer:                       customer,
		CustomerNotificationPreference: customerNotificationPreference,
		Currency:                       "IDR",
	}

	resp, err := invoice.Create(&data)
	fmt.Println(resp)
	if err != nil {
		log.Fatal(err)
	}
	var paymentUrl = resp.InvoiceURL
	hire.PaymentUrl = &paymentUrl

	err = huc.HireRepository.CreateHire(hire)

	return err

}

//func (huc *HireUseCase) AcceptHire(hire entities.Hire) error{
//	var hire =
//
//	hire, err := huc.HireRepository.
//}
