package weather

import "github.com/google/uuid"

type Weather struct {
	StadiumID   *uuid.UUID `json:"stadium_id,omitempty"`
	Temp        float64    `json:"temp"`
	Description string     `json:"description"`
	Icon        string     `json:"icon"`
}

type OpenWeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}
