package weather_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"us-soccer-go-test/internal/ent"
	"us-soccer-go-test/internal/ent/enttest"
	"us-soccer-go-test/internal/weather"

	"github.com/apex/log"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type weatherResponse struct {
	Weather weather.Weather `json:"weather"`
}

func mockHttpServerWithRouter(client *ent.Client) *httptest.Server {
	r := chi.NewRouter()
	r.Route("/api/weather", weather.NewController(log.Log, client).Route)

	ts := httptest.NewServer(r)

	return ts
}

func TestGetWeatherForStadium(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	testStadium := createTestStadium(client)

	ts1 := mockHttpServer(`{"coord":{"lon":-0.1086,"lat":51.555},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":22.34,"feels_like":21.91,"temp_min":20.44,"temp_max":23.58,"pressure":1027,"humidity":49},"visibility":10000,"wind":{"speed":4.12,"deg":340},"clouds":{"all":2},"dt":1717338248,"sys":{"type":2,"id":2075535,"country":"GB","sunrise":1717300069,"sunset":1717358968},"timezone":3600,"id":3333156,"name":"Islington","cod":200}`, 200)
	os.Setenv("OPENWEATHER_URL", fmt.Sprintf("%s/%s", ts1.URL, "?lat=%s&lon=%s&appid=%s&units=metric"))

	ts2 := mockHttpServerWithRouter(client)

	resp, err := http.Get(fmt.Sprintf("%s/api/weather/%s", ts2.URL, testStadium.ID.String()))

	if err != nil {
		t.Log("http get returned an err")
		t.FailNow()
	}

	if resp.StatusCode != 200 {
		t.Log("resp from mocked server with router returned non-200")
		t.FailNow()
	}

	defer resp.Body.Close()

	// decode the resp into a struct

	var decodedResponse weatherResponse

	err = json.NewDecoder(resp.Body).Decode(&decodedResponse)

	if err != nil {
		t.Log("json decode failed")
		t.FailNow()
	}

	result := decodedResponse.Weather

	if result.Temp != 22.34 {
		t.Logf("temp is wrong, expected 22.34, got %f instead", result.Temp)
		t.FailNow()
	}

	if result.Description != "clear sky" {
		t.Logf("description is wrong, expected 'clear sky', got %s instead", result.Description)
		t.FailNow()
	}

	if result.Icon != "01d" {
		t.Logf("icon is wrong, expected '01d' got %s instead", result.Icon)
		t.FailNow()
	}

	ts1.Close()
	ts2.Close()
}

func TestGetWeatherErrorBadUuid(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	ts := mockHttpServerWithRouter(client)

	resp, err := http.Get(fmt.Sprintf("%s/api/weather/test", ts.URL))

	if err != nil {
		t.Log("http get returned an err")
		t.FailNow()
	}

	if resp.StatusCode != 422 {
		t.Log("resp from mocked server with router returned non-422")
		t.FailNow()
	}

	defer resp.Body.Close()

	// decode the resp into a struct

	var decodedResponse map[string]string

	err = json.NewDecoder(resp.Body).Decode(&decodedResponse)

	if err != nil {
		t.Log("json decode failed")
		t.FailNow()
	}

	if decodedResponse["error"] != "invalid stadium ID" {
		t.Logf("unexpected error field value")
		t.FailNow()
	}

	ts.Close()
}

func TestGetWeatherErrorNonExistentStadiumId(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	ts := mockHttpServerWithRouter(client)

	resp, err := http.Get(fmt.Sprintf("%s/api/weather/%s", ts.URL, uuid.New().String()))

	if err != nil {
		t.Log("http get returned an err")
		t.FailNow()
	}

	if resp.StatusCode != 404 {
		t.Log("resp from mocked server with router returned non-404")
		t.FailNow()
	}

	defer resp.Body.Close()

	// decode the resp into a struct

	var decodedResponse map[string]string

	err = json.NewDecoder(resp.Body).Decode(&decodedResponse)

	if err != nil {
		t.Log("json decode failed")
		t.FailNow()
	}

	if decodedResponse["error"] != "stadium not found" {
		t.Logf("unexpected error field value")
		t.FailNow()
	}

	ts.Close()
}
