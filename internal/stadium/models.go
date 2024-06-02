package stadium

import (
	"us-soccer-go-test/internal/location"
	"us-soccer-go-test/internal/weather"

	"github.com/google/uuid"
)

type Stadium struct {
	ID       uuid.UUID         `json:"id"`
	Stadium  string            `json:"stadium"`
	Location location.Location `json:"location"`
	Weather  *weather.Weather  `json:"weather,omitempty"`
}
