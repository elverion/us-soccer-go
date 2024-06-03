package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"us-soccer-go-test/internal/ent"
	"us-soccer-go-test/internal/ent/stadium"
	"us-soccer-go-test/internal/ent/weather"

	"github.com/apex/log"
	"github.com/google/uuid"
)

const baseURL = "https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric"

func (c *Controller) getWeatherByStadium(stadiumID string, includeID bool, ctx context.Context) (*Weather, error) {
	id, _ := uuid.Parse(stadiumID)
	stadium, err := c.db.Stadium.Query().Where(stadium.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	w, _ := FetchWeather(strconv.FormatFloat(stadium.Latitude, 'f', -1, 64), strconv.FormatFloat(stadium.Longitude, 'f', -1, 64), c.db, stadium.ID, ctx, c.logger)

	if includeID {
		return &Weather{StadiumID: &id, Description: w.Description, Icon: w.Icon, Temp: w.Temp}, nil
	}

	return w, nil
}

func FetchWeather(lat, long string, db *ent.Client, stadiumId uuid.UUID, ctx context.Context, logger log.Interface) (*Weather, bool) {
	// Check the DB for a weather entry for the stadium
	result, err := db.Weather.Query().Where(weather.HasStadiumWith(stadium.ID(stadiumId))).First(ctx)

	var new bool

	if err != nil {
		if !ent.IsNotFound(err) {
			// log an error
			logger.WithError(err).Warn("failed to query db")
			return nil, false
		}
		new = true
	} else {
		if time.Since(result.UpdateTime).Minutes() < 10 {
			// don't do anything
			return &Weather{
				Temp:        result.Temperature,
				Description: result.Description,
				Icon:        result.Icon,
			}, false
		}
		new = false
	}

	resp, err := http.Get(fmt.Sprintf(baseURL, lat, long, os.Getenv("API_KEYS_OPENWEATHER")))

	if err != nil {
		// handle error
		logger.WithError(err).Warn("Failed to query openweather API")
		if new {
			return nil, false
		}
		return &Weather{
			Temp:        result.Temperature,
			Description: result.Description,
			Icon:        result.Icon,
		}, false
	}

	if resp.StatusCode != 200 {
		// handle non 200 status code
		logger.Warn("Got non-200 back from openweather API")
		if new {
			return nil, false
		}

		return &Weather{
			Temp:        result.Temperature,
			Description: result.Description,
			Icon:        result.Icon,
		}, false
	}

	defer resp.Body.Close()

	var owr OpenWeatherResponse // Value is not set yet so it doesn't violate functional programming

	json.NewDecoder(resp.Body).Decode(&owr)

	// Create (or update) an entry in the database for the weather

	var newResult *ent.Weather

	if new {
		newResult, err = db.Weather.Create().SetDescription(owr.Weather[0].Description).SetTemperature(owr.Main.Temp).SetIcon(owr.Weather[0].Icon).SetStadiumID(stadiumId).Save(ctx)
	} else {
		newResult, err = result.Update().SetDescription(owr.Weather[0].Description).SetIcon(owr.Weather[0].Icon).SetTemperature(owr.Main.Temp).Save(ctx)
	}

	if err != nil {
		logger.WithError(err).Warn("Failed to upsert weather data")
		if new {
			return nil, true
		}

		return &Weather{
			Temp:        result.Temperature,
			Description: result.Description,
			Icon:        result.Icon,
		}, true
	}

	return &Weather{
		Temp:        newResult.Temperature,
		Description: newResult.Description,
		Icon:        newResult.Icon,
	}, true

}
