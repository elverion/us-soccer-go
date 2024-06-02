package stadium

import "github.com/go-chi/chi/v5"

func (c *Controller) Route(r chi.Router) {
	r.Post("/", c.upload)
	r.Get("/", c.getAllStadiums)
}
