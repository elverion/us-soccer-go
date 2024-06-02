package weather

import (
	"net/http"
	"us-soccer-go-test/internal/ent"

	"github.com/apex/log"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/lrstanley/chix"
)

type Controller struct {
	logger log.Interface
	db     *ent.Client
}

func NewController(lg log.Interface, db *ent.Client) *Controller {
	return &Controller{logger: lg, db: db}
}

func (c *Controller) getWeatherForAllStadiums(w http.ResponseWriter, r *http.Request) {
	weather, err := c.getWeatherForStadiums(r.Context())
	if chix.Error(w, r, err) {
		return
	}

	chix.JSON(w, r, 200, chix.M{"weather": weather})
}

func (c *Controller) getWeatherForStadium(w http.ResponseWriter, r *http.Request) {
	stadiumID := chi.URLParam(r, "stadium")

	if stadiumID == "" {
		chix.JSON(w, r, 422, chix.M{"error": "missing stadium ID"})
		return
	}

	_, err := uuid.Parse(stadiumID)

	if err != nil {
		chix.JSON(w, r, 422, chix.M{"error": "invalid stadium ID"})
		return
	}

	weather, err := c.getWeatherByStadium(stadiumID, false, r.Context())

	if err != nil {
		if !ent.IsNotFound(err) {
			chix.Error(w, r, err)
			return
		}

		chix.JSON(w, r, 404, chix.M{"error": "stadium not found"})
		return
	}

	chix.JSON(w, r, 200, chix.M{"weather": weather})
}
