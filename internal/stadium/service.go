package stadium

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"us-soccer-go-test/internal/ent"
	"us-soccer-go-test/internal/location"
	"us-soccer-go-test/internal/weather"
)

func (c *Controller) getLatLongFloats(lat, long string) (float64, float64) {
	l1, _ := strconv.ParseFloat(lat, 32)
	l2, _ := strconv.ParseFloat(long, 32)

	return l1, l2
}

func (c *Controller) validateCSV(index int, records [][]string) (int, bool) {
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

	if !isFloat(line[LatIndex]) || !isFloat(line[LongIndex]) {
		return -1, false
	}

	isEmpty := func(input string) bool {
		trimmedInput := strings.TrimSpace(input)

		return trimmedInput == ""
	}

	// The rest of the columns must not be empty
	if isEmpty(line[TeamIndex]) || isEmpty(line[FDCOUKIndex]) || isEmpty(line[CityIndex]) || isEmpty(line[StadiumIndex]) || isEmpty(line[CapacityIndex]) || isEmpty(line[CountryIndex]) {
		return -1, false
	}

	return c.validateCSV(index+1, records)
}

func (c *Controller) insertRecord(index int, records [][]string, ctx context.Context) (int, bool) {
	if index > len(records)-1 || index == -1 {
		// -1 is failure
		return len(records) - 1, index != -1
	}

	line := records[index]

	capacity, err := strconv.Atoi(line[CapacityIndex])
	lat, long := c.getLatLongFloats(line[LatIndex], line[LongIndex])

	if err != nil {
		return -1, false
	}

	_, err = c.db.Stadium.Create().
		SetTeam(strings.TrimSpace(line[TeamIndex])).
		SetFdcouk(strings.TrimSpace(line[FDCOUKIndex])).
		SetStadium(strings.TrimSpace(line[StadiumIndex])).
		SetCapacity(capacity).
		SetCity(strings.TrimSpace(line[CityIndex])).
		SetLatitude(lat).
		SetLongitude(long).
		SetCountry(strings.TrimSpace(line[CountryIndex])).Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			// likely already exists, move ot next
			return c.insertRecord(index+1, records, ctx)
		}
		c.logger.WithError(err).Warn("Failed to save stadium")
		return -1, false
	}

	return c.insertRecord(index+1, records, ctx)
}

func (c *Controller) handleCSV(r *http.Request) error {
	formFile, _, err := r.FormFile("csv")

	if err != nil {
		return err
	}

	// Read the CSV
	reader := csv.NewReader(formFile)

	records, err := reader.ReadAll()

	if err != nil {
		return err
	}

	formFile.Close()

	// validate the CSV
	_, valid := c.validateCSV(1, records)
	if !valid {
		return fmt.Errorf("invalid CSV")
	}

	_, success := c.insertRecord(1, records, r.Context())

	if !success {
		return fmt.Errorf("db error")
	}

	return nil
}

func (c *Controller) getStadiumsFunc(index int, records []*ent.Stadium, data []Stadium, withWeather bool) []Stadium {
	if index > len(records)-1 {
		return data
	}

	record := records[index]

	var w *weather.Weather
	if withWeather {
		if record.Edges.Weather == nil {
			w = &weather.Weather{
				Description: "unavailable",
				Icon:        "01d",
			}
		} else {
			w = &weather.Weather{
				Temp:        record.Edges.Weather.Temperature,
				Description: record.Edges.Weather.Description,
				Icon:        record.Edges.Weather.Icon,
			}
		}
	}

	stadium := Stadium{
		ID:      record.ID,
		Stadium: record.Stadium,
		Location: location.Location{
			City:    record.City,
			Country: record.Country,
			Lat:     record.Latitude,
			Long:    record.Longitude,
		},

		Weather: w,
	}

	newData := append(data, stadium)

	return c.getStadiumsFunc(index+1, records, newData, withWeather)
}

func (c *Controller) getStadiums(ctx context.Context, withWeather bool) ([]Stadium, error) {
	// get the stadiums from the db

	query := c.db.Stadium.Query()

	if withWeather {
		query.WithWeather()
	}
	records, err := query.All(ctx)

	if err != nil {
		return nil, err
	}

	return c.getStadiumsFunc(0, records, []Stadium{}, withWeather), nil
}
