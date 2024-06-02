package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	weather, err := c.db.Weather.Query().Where(weather.HasStadiumWith(stadium.ID(id))).First(ctx)
	if err != nil {
		return nil, err
	}

	if includeID {
		return &Weather{StadiumID: &id, Description: weather.Description, Temp: weather.Temperature}, nil
	}

	return &Weather{Description: weather.Description, Temp: weather.Temperature}, nil
}

func (c *Controller) getWeatherForStadiumsFunc(ctx context.Context, index int, records []*ent.Stadium, weather []Weather) (int, []Weather) {
	if index > len(records)-1 {
		return -1, weather
	}

	w, err := c.getWeatherByStadium(string(records[index].ID.String()), true, ctx)

	if err != nil {
		c.logger.WithError(err).Warn("Could not get weather")
		return -1, nil
	}

	newData := append(weather, *w)

	return c.getWeatherForStadiumsFunc(ctx, index+1, records, newData)
}

func (c *Controller) getWeatherForStadiums(ctx context.Context) ([]Weather, error) {
	stadiums, err := c.db.Stadium.Query().All(ctx)

	if err != nil {
		return nil, err
	}

	_, weather := c.getWeatherForStadiumsFunc(ctx, 0, stadiums, []Weather{})

	return weather, nil
}

func fetchWeather(lat, long string, db *ent.Client, stadiumId uuid.UUID, ctx context.Context) {
	// Check the DB for a weather entry for the stadium
	result, err := db.Weather.Query().Where(weather.HasStadiumWith(stadium.ID(stadiumId))).First(ctx)

	var new bool

	if err != nil {
		if !ent.IsNotFound(err) {
			// log an error
			log.WithError(err).Warn("failed to query db")
			return
		}
		new = true
	} else {
		if time.Since(result.UpdateTime).Minutes() < 10 {
			// don't do anything
			return
		}
		new = false
	}

	resp, err := http.Get(fmt.Sprintf(baseURL, lat, long, os.Getenv("API_KEYS_OPENWEATHER")))

	if err != nil {
		// handle error
		log.WithError(err).Warn("Failed to query openweather API")
		return
	}

	if resp.StatusCode != 200 {
		// handle non 200 status code
		log.Warn("Got non-200 back from openweather API")
		return
	}

	defer resp.Body.Close()

	var owr OpenWeatherResponse // Value is not set yet so it doesn't violate functional programming

	json.NewDecoder(resp.Body).Decode(&owr)

	// Create (or update) an entry in the database for the weather

	if new {
		db.Weather.Create().SetDescription(owr.Weather[0].Description).SetTemperature(owr.Main.Temp).SetStadiumID(stadiumId).Save(ctx)
	} else {
		result.Update().SetDescription(owr.Weather[0].Description).SetTemperature(owr.Main.Temp).Save(ctx)
	}
}
