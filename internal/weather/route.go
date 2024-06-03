package weather

import "github.com/go-chi/chi/v5"

func (c *Controller) Route(r chi.Router) {
	r.Get("/{stadium}", c.getWeatherForStadium)
}
