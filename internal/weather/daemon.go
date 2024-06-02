package weather

import (
	"context"
	"strconv"
	"time"
	"us-soccer-go-test/internal/ent"

	"github.com/apex/log"
)

func StartWeatherDaemon(db *ent.Client) func() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(db *ent.Client, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				stadiums, err := db.Stadium.Query().All(ctx)
				if err != nil {
					log.WithError(err).Warn("failed to fetch locations")
				} else {
					count := 0
					for _, stadium := range stadiums {
						fetchWeather(strconv.FormatFloat(stadium.Latitude, 'f', -1, 64), strconv.FormatFloat(stadium.Longitude, 'f', -1, 64), db, stadium.ID, ctx)

						if count >= 59 {
							// sleep for a minute to reset the API key count
							time.Sleep(1 * time.Minute)
							count = 0
						}
						count++
					}
				}
				time.Sleep(1 * time.Minute)
			}
		}

	}(db, ctx)

	return cancel

}
