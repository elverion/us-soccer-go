package weather_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
	"us-soccer-go-test/internal/ent/enttest"
	"us-soccer-go-test/internal/weather"

	"github.com/apex/log"
)

func TestDaemon(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Set up a mock HTTP server
	ts := mockHttpServer(`{"coord":{"lon":-0.1086,"lat":51.555},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":22.34,"feels_like":21.91,"temp_min":20.44,"temp_max":23.58,"pressure":1027,"humidity":49},"visibility":10000,"wind":{"speed":4.12,"deg":340},"clouds":{"all":2},"dt":1717338248,"sys":{"type":2,"id":2075535,"country":"GB","sunrise":1717300069,"sunset":1717358968},"timezone":3600,"id":3333156,"name":"Islington","cod":200}`, 200)

	createTestStadium(client)

	os.Setenv("OPENWEATHER_URL", fmt.Sprintf("%s/%s", ts.URL, "?lat=%s&lon=%s&appid=%s&units=metric"))

	cancel := weather.StartWeatherDaemon(client, log.Log)

	// wait a few seconds

	time.Sleep(2 * time.Second)

	// now check the database

	count, err := client.Weather.Query().Count(context.Background())

	if err != nil {
		t.FailNow()
	}

	if count != 1 {
		t.Logf("Expected count of 1, got %d instead", count)
		t.FailNow()
	}
	cancel()
	ts.Close()
}
