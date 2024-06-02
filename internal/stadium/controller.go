package stadium

import (
	"net/http"
	"us-soccer-go-test/internal/ent"

	"github.com/apex/log"
	"github.com/lrstanley/chix"
)

type Controller struct {
	logger log.Interface
	db     *ent.Client
}

func NewController(lg log.Interface, db *ent.Client) *Controller {
	return &Controller{logger: lg, db: db}
}

func (c *Controller) upload(w http.ResponseWriter, r *http.Request) {
	// Limits to 1 MB
	if chix.Error(w, r, r.ParseMultipartForm(1<<20)) {
		return
	}

	err := c.handleCSV(r)

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

func (c *Controller) getAllStadiums(w http.ResponseWriter, r *http.Request) {
	var includeWeather bool

	if r.URL.Query().Get("weather") == "1" {
		includeWeather = true
	} else {
		includeWeather = false
	}

	stadiums, err := c.getStadiums(r.Context(), includeWeather)

	if chix.Error(w, r, err) {
		return
	}

	chix.JSON(w, r, 200, chix.M{"stadiums": stadiums})
}
