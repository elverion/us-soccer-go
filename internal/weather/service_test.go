package weather_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"us-soccer-go-test/internal/ent"
	"us-soccer-go-test/internal/ent/enttest"
	"us-soccer-go-test/internal/weather"

	_ "github.com/mattn/go-sqlite3"

	"github.com/apex/log"
)

func mockHttpServer(response string, status int) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprintln(w, response)
	}))

	return ts
}

func createTestStadium(client *ent.Client) *ent.Stadium {
	testStadium, _ := client.Stadium.Create().SetCapacity(4).SetCountry("Test").SetFdcouk("test").SetCity("test").SetLatitude(4).SetLongitude(4).SetTeam("test").SetStadium("test").Save(context.Background())

	return testStadium
}

func TestFetchWeatherSuccessful(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Set up a mock HTTP server
	ts := mockHttpServer(`{"coord":{"lon":-0.1086,"lat":51.555},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":22.34,"feels_like":21.91,"temp_min":20.44,"temp_max":23.58,"pressure":1027,"humidity":49},"visibility":10000,"wind":{"speed":4.12,"deg":340},"clouds":{"all":2},"dt":1717338248,"sys":{"type":2,"id":2075535,"country":"GB","sunrise":1717300069,"sunset":1717358968},"timezone":3600,"id":3333156,"name":"Islington","cod":200}`, 200)

	testStadium := createTestStadium(client)

	os.Setenv("OPENWEATHER_URL", fmt.Sprintf("%s/%s", ts.URL, "?lat=%s&lon=%s&appid=%s&units=metric"))

	result, _ := weather.FetchWeather("1", "2", client, testStadium.ID, context.Background(), log.Log)

	if result.Temp != 22.34 {
		t.FailNow()
	}

	if result.Description != "clear sky" {
		t.FailNow()
	}

	if result.Icon != "01d" {
		t.FailNow()
	}

	ts.Close()
}

func TestFetchWeatherCache(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Set up a mock HTTP server
	ts := mockHttpServer(`{"coord":{"lon":-0.1086,"lat":51.555},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":22.34,"feels_like":21.91,"temp_min":20.44,"temp_max":23.58,"pressure":1027,"humidity":49},"visibility":10000,"wind":{"speed":4.12,"deg":340},"clouds":{"all":2},"dt":1717338248,"sys":{"type":2,"id":2075535,"country":"GB","sunrise":1717300069,"sunset":1717358968},"timezone":3600,"id":3333156,"name":"Islington","cod":200}`, 200)

	testStadium := createTestStadium(client)

	os.Setenv("OPENWEATHER_URL", fmt.Sprintf("%s/%s", ts.URL, "?lat=%s&lon=%s&appid=%s&units=metric"))

	weather.FetchWeather("1", "2", client, testStadium.ID, context.Background(), log.Log)

	ts.Close()

	// recreate the tester server, but change the temperature

	ts = mockHttpServer(`{"coord":{"lon":-0.1086,"lat":51.555},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":30.21,"feels_like":21.91,"temp_min":20.44,"temp_max":23.58,"pressure":1027,"humidity":49},"visibility":10000,"wind":{"speed":4.12,"deg":340},"clouds":{"all":2},"dt":1717338248,"sys":{"type":2,"id":2075535,"country":"GB","sunrise":1717300069,"sunset":1717358968},"timezone":3600,"id":3333156,"name":"Islington","cod":200}`, 200)

	result, _ := weather.FetchWeather("1", "2", client, testStadium.ID, context.Background(), log.Log)

	// the temp should not be equal to 30.21
	if result.Temp == 30.21 {
		t.FailNow()
	}

	// it should still be 22.34
	if result.Temp != 22.34 {
		t.FailNow()
	}

	ts.Close()
}

func TestFetchWeatherAPIFail(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Set up a mock HTTP server
	ts := mockHttpServer(`{"cod":429, "message":"Too many requests"}`, 429)

	testStadium := createTestStadium(client)

	os.Setenv("OPENWEATHER_URL", fmt.Sprintf("%s/%s", ts.URL, "?lat=%s&lon=%s&appid=%s&units=metric"))

	result, _ := weather.FetchWeather("1", "2", client, testStadium.ID, context.Background(), log.Log)

	if result != nil {
		t.FailNow()
	}

	ts.Close()
}
