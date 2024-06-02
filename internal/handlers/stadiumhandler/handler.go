package stadiumhandler

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"us-soccer-go-test/internal/models"
	"us-soccer-go-test/internal/utils"

	"github.com/apex/log"
	"github.com/go-chi/chi/v5"
	"github.com/lrstanley/chix"
)

const teamIndex = 0
const fdcoukIndex = 1
const cityIndex = 2
const stadiumIndex = 3
const capacityIndex = 4
const latIndex = 5
const longIndex = 6
const countryIndex = 7

func validateFunc(index int, records [][]string) (int, bool) {
	if index > len(records)-1 || index == -1 {
		// -1 is failure
		return len(records) - 1, index != -1
	}

	line := records[index]

	if len(line) != 8 {
		return -1, false
	}

	// columns 5 and 6, must be floats

	isFloat := func(input string) bool {
		_, err := strconv.ParseFloat(input, 32)
		return err == nil
	}

	if !isFloat(line[latIndex]) || !isFloat(line[longIndex]) {
		return -1, false
	}

	isEmpty := func(input string) bool {
		trimmedInput := strings.TrimSpace(input)

		return trimmedInput == ""
	}

	// The rest of the columns must not be empty
	if isEmpty(line[teamIndex]) || isEmpty(line[fdcoukIndex]) || isEmpty(line[cityIndex]) || isEmpty(line[stadiumIndex]) || isEmpty(line[capacityIndex]) || isEmpty(line[countryIndex]) {
		return -1, false
	}

	return validateFunc(index+1, records)
}

func validateCSV(file multipart.File) bool {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		return false
	}

	_, valid := validateFunc(1, records)

	return valid
}

func createStadiumFunc(index int, records [][]string, data []models.Stadium) []models.Stadium {
	if index > len(records)-1 {
		return data
	}

	line := records[index]

	lat, long := func(lat string, long string) (float64, float64) {
		l1, _ := strconv.ParseFloat(lat, 32)
		l2, _ := strconv.ParseFloat(long, 32)

		return l1, l2
	}(line[5], line[6])

	stadium := models.Stadium{
		Stadium: strings.TrimSpace(line[3]),
		Location: models.Location{
			City:    strings.TrimSpace(line[2]),
			Country: strings.TrimSpace(line[7]),
			Lat:     lat,
			Long:    long,
		},

		Weather: utils.FetchWeather(line[5], line[6]),
	}

	newData := append(data, stadium)

	return createStadiumFunc(index+1, records, newData)
}

type Handler struct {
	logger log.Interface
}

func NewHandler(lg log.Interface) *Handler {
	return &Handler{logger: lg}
}

func (h *Handler) Route(r chi.Router) {
	r.Post("/", h.upload)
	r.Get("/", h.getAllStadiums)
}

func (h *Handler) upload(w http.ResponseWriter, r *http.Request) {
	// Limits to 1 MB
	if chix.Error(w, r, r.ParseMultipartForm(1<<20)) {
		return
	}

	err := h.handleCSV(r)

	if err != nil {
		if err.Error() == "invalid CSV" {
			chix.JSON(w, r, 400, chix.M{"success": false, "error": err.Error()})
			return
		}

		chix.Error(w, r, err)
		return
	}

	chix.JSON(w, r, 200, chix.M{"success": true})

}

func (h *Handler) getAllStadiums(w http.ResponseWriter, r *http.Request) {
	// open the CSV file from disk

	file, err := os.Open("./storage/stadiums.csv")

	if chix.Error(w, r, err) {
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if chix.Error(w, r, err) {
		return
	}

	stadiums := createStadiumFunc(1, records, []models.Stadium{})

	chix.JSON(w, r, 200, chix.M{"stadiums": stadiums})
}

func (h *Handler) handleCSV(r *http.Request) error {
	formFile, _, err := r.FormFile("csv")

	if err != nil {
		return err
	}

	// validate the CSV
	if !validateCSV(formFile) {
		return fmt.Errorf("invalid CSV")
	}

	// reset
	formFile.Seek(0, io.SeekStart)

	file, err := os.Create("./storage/stadiums.csv")

	if err != nil {
		return err
	}
	_, err = io.Copy(file, formFile)

	if err != nil {
		return err
	}

	file.Close()
	formFile.Close()

	return nil
}
