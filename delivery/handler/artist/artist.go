package artist

import (
	"net/http"
	"potentivio-app/delivery/helper"
	_middlewares "potentivio-app/delivery/middlewares"
	"potentivio-app/entities"
	_artistUseCase "potentivio-app/usecase/artist"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArtistHandler struct {
	artistUseCase _artistUseCase.ArtistUseCaseInterface
}

func NewArtistHandler(artistUseCase _artistUseCase.ArtistUseCaseInterface) *ArtistHandler {
	return &ArtistHandler{
		artistUseCase: artistUseCase,
	}
}

func (ah *ArtistHandler) CreateArtistHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newArtist entities.Artist
		err := c.Bind(&newArtist)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}
		_, error := ah.artistUseCase.CreateArtist(newArtist)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to create user"))
	}
}

func (ah *ArtistHandler) GetAllArtistHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		id_catagory, _ := strconv.Atoi(c.QueryParam("id_catagory"))
		id_genre, _ := strconv.Atoi(c.QueryParam("id_genre"))

		filters_catagory_genre := map[string]int{}
		if id_catagory != 0 {
			filters_catagory_genre["id_catagory"] = id_catagory
		}
		if id_genre != 0 {
			filters_catagory_genre["id_genre"] = id_genre
		}

		filters_address := map[string]string{}
		if c.QueryParam("address") != "" {
			filters_address["address"] = c.QueryParam("address")
		}
		if c.QueryParam("name") != "" {
			filters_address["name"] = c.QueryParam("name")
		}

		filters_price := map[string]string{}
		if c.QueryParam("price") != "" {
			filters_price["price"] = c.QueryParam("price")
		}

		artists, rows, err := ah.artistUseCase.GetAllArtist(filters_catagory_genre, filters_price, filters_address)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		responseArtist := []map[string]interface{}{}
		for i := 0; i < len(artists); i++ {
			response := map[string]interface{}{
				"id":          artists[i].ID,
				"id_catagory": artists[i].IdCatagory,
				"id_genre":    artists[i].IdGenre,
				"artist_name": artists[i].Name,
				"address":     artists[i].Address,
				"price":       artists[i].Price,
				"description": artists[i].Description,
				"avatar":      artists[i].Avatar,
				"catagory": map[string]interface{}{
					"id":            artists[i].Catagory.ID,
					"name_catagory": artists[i].Catagory.NameCatagory,
				},
				"genre": map[string]interface{}{
					"id":         artists[i].Genre.ID,
					"name_genre": artists[i].Genre.NameGenre,
				},
			}
			responseArtist = append(responseArtist, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get all artist", responseArtist))
	}
}

func (ah *ArtistHandler) GetProfileArtistHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		//mendapatkan id dari token yang login
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		artist, rows, err := ah.artistUseCase.GetProfileArtist(uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		video_artist := []map[string]interface{}{}
		for i := 0; i < len(artist.VideoArtist); i++ {
			response := map[string]interface{}{
				"id":        artist.VideoArtist[i].ID,
				"video_url": artist.VideoArtist[i].VideoUrl,
			}
			video_artist = append(video_artist, response)
		}

		total, rating, errCount := ah.artistUseCase.CountRating(artist.ID)
		artist.TotalRate = total
		artist.Rating = rating
		if errCount != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errCount.Error()))
		}

		responseArtist := map[string]interface{}{
			"id":             artist.ID,
			"artist_name":    artist.Name,
			"email":          artist.Email,
			"rating":         artist.Rating,
			"total_rate":     artist.TotalRate,
			"id_catagory":    artist.IdCatagory,
			"id_genre":       artist.IdGenre,
			"phone_number":   artist.PhoneNumber,
			"address":        artist.Address,
			"price":          artist.Price,
			"description":    artist.Description,
			"account_number": artist.AccountNumber,
			"avatar":         artist.Avatar,
			"name_catagory":  artist.Catagory.NameCatagory,
			"name_genre":     artist.Genre.NameGenre,
			"video_artist":   video_artist,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get profile artist", responseArtist))
	}
}

func (ah *ArtistHandler) GetArtistByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		artist, hires, rows, err := ah.artistUseCase.GetArtistById(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		notAvailable := []map[string]interface{}{}
		history := []map[string]interface{}{}

		for i := 0; i < len(hires); i++ {
			if hires[i].StatusCafe == "waiting payment" {
				response := map[string]interface{}{
					"cafe_name": hires[i].Cafe.Name,
					"date":      hires[i].Date,
				}
				notAvailable = append(notAvailable, response)
			}
			if hires[i].StatusCafe == "done" {
				response := map[string]interface{}{
					"cafe_name": hires[i].Cafe.Name,
					"date":      hires[i].Date,
					"rating":    hires[i].Rating,
					"comment":   hires[i].Comment,
				}
				history = append(history, response)
			}
		}

		videoArtist := []map[string]interface{}{}
		for i := 0; i < len(artist.VideoArtist); i++ {
			response := map[string]interface{}{
				"id":        artist.VideoArtist[i].ID,
				"video_url": artist.VideoArtist[i].VideoUrl,
			}
			videoArtist = append(videoArtist, response)
		}

		total, rating, errCount := ah.artistUseCase.CountRating(artist.ID)
		artist.TotalRate = total
		artist.Rating = rating
		if errCount != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errCount.Error()))
		}

		responseArtist := map[string]interface{}{
			"id":             artist.ID,
			"rating":         artist.Rating,
			"total_rate":     artist.TotalRate,
			"artist_name":    artist.Name,
			"id_catagory":    artist.IdCatagory,
			"id_genre":       artist.IdGenre,
			"phone_number":   artist.PhoneNumber,
			"address":        artist.Address,
			"price":          artist.Price,
			"description":    artist.Description,
			"account_number": artist.AccountNumber,
			"avatar":         artist.Avatar,
			"video_artist":   videoArtist,
			"not_available":  notAvailable,
			"hire_history":   history,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get detail artist", responseArtist))

	}
}

func (ah *ArtistHandler) UpdateArtistHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang login
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var updateArtist entities.Artist
		errBind := c.Bind(&updateArtist)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errBind.Error()))
		}

		// prosess binding image
		fileData, fileInfo, err_binding_image := c.Request().FormFile("avatar")
		if err_binding_image != http.ErrMissingFile {
			if err_binding_image != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFailed("bind image error"))
			}

			// check file extension
			_, err_check_extension := helper.CheckFileExtension(fileInfo.Filename)
			if err_check_extension != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file extension error"))
			}

			// check file size
			err_check_size := helper.CheckFileSize(fileInfo.Size)
			if err_check_size != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file size error"))
			}

			// memberikan nama file
			fileName := "foto_profile_" + strconv.Itoa(idToken)

			// upload foto profile
			var err_upload_photo error
			theUrl, err_upload_photo := helper.UploadImage("foto_profile_artist", fileName, fileData)
			if err_upload_photo != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFailed("upload image failed"))
			}

			// create foto profile artist
			updateArtist.Avatar = &theUrl
		}

		_, rows, err := ah.artistUseCase.UpdateArtist(updateArtist, uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to update artist"))
	}
}

func (ah *ArtistHandler) DeleteArtistHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		//mendapatkan id dari token yang login
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		rows, err := ah.artistUseCase.DeleteArtist(uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to delete artist"))
	}
}
