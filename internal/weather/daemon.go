package weather

import (
	"context"
	"strconv"
	"time"
	"us-soccer-go-test/internal/ent"

	"github.com/apex/log"
)

func StartWeatherDaemon(db *ent.Client, logger log.Interface) func() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(db *ent.Client, ctx context.Context, logger log.Interface) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				stadiums, err := db.Stadium.Query().All(ctx)
				if err != nil {
					logger.WithError(err).Warn("failed to fetch stadiums")
				} else {
					count := 0
					updateCount := 0
					for _, stadium := range stadiums {
						_, hitAPI := FetchWeather(strconv.FormatFloat(stadium.Latitude, 'f', -1, 64), strconv.FormatFloat(stadium.Longitude, 'f', -1, 64), db, stadium.ID, ctx, logger)

						// if it didn't hit the API successfully or if it returned a cache from the DB, no need to increment the count.
						if !hitAPI {
							continue
						}

						if count >= 59 {
							// sleep for a minute to reset the API key count
							time.Sleep(1 * time.Minute)
							count = 0
						}
						count++
						updateCount++
					}
					if updateCount > 0 {
						logger.Infof("Updated weather data for %d stadiums", updateCount)
					}
				}
				time.Sleep(1 * time.Minute)
			}
		}

	}(db, ctx, logger)

	return cancel

}
