package stadium_test

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"sort"
	"strconv"
	"strings"
	"testing"
	"us-soccer-go-test/internal/ent"
	"us-soccer-go-test/internal/ent/enttest"
	"us-soccer-go-test/internal/stadium"
	"us-soccer-go-test/internal/weather"

	"github.com/apex/log"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

type stadiumResponse struct {
	Stadiums []stadium.Stadium `json:"stadiums"`
}

const csvData = `Team,FDCOUK,City,Stadium,Capacity,Latitude,Longitude,Country
Arsenal ,Arsenal,London ,Emirates Stadium ,60361,51.555,-0.108611,England
Aston Villa ,Aston Villa,Birmingham ,Villa Park ,42785,52.509167,-1.884722,England`

func mockHttpServerWithRouter(client *ent.Client) *httptest.Server {
	r := chi.NewRouter()
	r.Route("/api/stadiums", stadium.NewController(log.Log, client).Route)

	ts := httptest.NewServer(r)

	return ts
}

func ingestCsv(url string, overriddenCsv, overriddenFieldName string) error {

	body := &bytes.Buffer{}

	var bydData []byte

	fieldName := "csv"

	if overriddenCsv != "" {
		bydData = []byte(overriddenCsv)
	} else {
		bydData = []byte(csvData)
	}

	if overriddenFieldName != "" {
		fieldName = overriddenFieldName
	}

	bytReader := bytes.NewReader(bydData)

	wrt := multipart.NewWriter(body)

	field, err := wrt.CreateFormFile(fieldName, "stadiums.csv")

	if err != nil {
		return err
	}

	io.Copy(field, bytReader)
	wrt.Close()

	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", wrt.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("expected 200 got %d instead", resp.StatusCode)
	}

	return nil
}

func compareCsvToStadium(line []string, stadia stadium.Stadium) bool {
	if strings.TrimSpace(line[stadium.StadiumIndex]) != stadia.Stadium {
		return false
	}

	if strings.TrimSpace(line[stadium.CityIndex]) != stadia.Location.City {
		return false
	}

	if strings.TrimSpace(line[stadium.CountryIndex]) != stadia.Location.Country {
		return false
	}

	lat, err := strconv.ParseFloat(line[stadium.LatIndex], 32)

	if err != nil {
		return false
	}

	long, err := strconv.ParseFloat(line[stadium.LongIndex], 32)

	if err != nil {
		return false
	}

	if lat != stadia.Location.Lat {
		return false
	}

	if long != stadia.Location.Long {
		return false
	}

	return true
}

func TestIngestCsvRoute(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	ts := mockHttpServerWithRouter(client)

	err := ingestCsv(fmt.Sprintf("%s/api/stadiums", ts.URL), "", "")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	ts.Close()
}

func TestIngestCsvRouteWithBadData(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	ts := mockHttpServerWithRouter(client)

	err := ingestCsv(fmt.Sprintf("%s/api/stadiums", ts.URL), fmt.Sprintf("%s,extra", csvData), "")

	if err == nil {
		t.FailNow()
	}

	err = ingestCsv(fmt.Sprintf("%s/api/stadiums", ts.URL), fmt.Sprintf("%s\nAston Villa ,Aston Villa,Birmingham ,Villa Park ,42785,bad.509167,-1.884722,England", csvData), "")

	if err == nil {
		t.FailNow()
	}

	err = ingestCsv(fmt.Sprintf("%s/api/stadiums", ts.URL), fmt.Sprintf("%s\n,Aston Villa,Birmingham ,Villa Park ,42785,52.509167,-1.884722,England", csvData), "")

	if err == nil {
		t.FailNow()
	}

	err = ingestCsv(fmt.Sprintf("%s/api/stadiums", ts.URL), `test,test2,test3
	a, b, c`, "")

	if err == nil {
		t.FailNow()
	}

	ts.Close()
}

func TestIngestCsvRouteWithDuplicates(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	ts := mockHttpServerWithRouter(client)

	err := ingestCsv(fmt.Sprintf("%s/api/stadiums", ts.URL), "", "")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	// re-run with same data
	err = ingestCsv(fmt.Sprintf("%s/api/stadiums", ts.URL), "", "")
	// duplicates should not cause an error to be returned from the API
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	count, err := client.Stadium.Query().Count(context.Background())

	if err != nil {
		t.FailNow()
	}

	// count should only equal 2

	if count != 2 {
		t.Logf("Expected count of 2, got %d instead", count)
		t.FailNow()
	}

	ts.Close()
}

func TestIngestCsvRouteWithoutFile(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	ts := mockHttpServerWithRouter(client)

	resp, err := http.Post(fmt.Sprintf("%s/api/stadiums", ts.URL), "text/csv", nil)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	// Chix defaults to setting the status code of 500 for errors returned that don't have an error class setup
	if resp.StatusCode != 500 {
		t.FailNow()
	}

	ts.Close()
}

func TestIngestCsvRouteWithoutWrongFieldName(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	ts := mockHttpServerWithRouter(client)

	err := ingestCsv(fmt.Sprintf("%s/api/stadiums", ts.URL), "", "file")

	if err == nil {
		t.FailNow()
	}

	ts.Close()
}

func TestGetStadiums(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	ts := mockHttpServerWithRouter(client)
	// ingest our data
	err := ingestCsv(fmt.Sprintf("%s/api/stadiums", ts.URL), "", "")

	if err != nil {
		t.FailNow()
	}

	// Make a call to get our stadiums now

	resp, err := http.Get(fmt.Sprintf("%s/api/stadiums", ts.URL))

	if err != nil {
		t.FailNow()
	}

	var decodedResponse stadiumResponse

	err = json.NewDecoder(resp.Body).Decode(&decodedResponse)

	if err != nil {
		t.Log("json decode failed")
		t.FailNow()
	}

	count := len(decodedResponse.Stadiums)

	if count != 2 {
		t.Logf("Expected count of 2, got %d instead", count)
		t.FailNow()
	}
	bytReader := bytes.NewReader([]byte(csvData))
	csvReader := csv.NewReader(bytReader)

	lines, err := csvReader.ReadAll()
	if err != nil {
		t.FailNow()
	}

	stadiums := decodedResponse.Stadiums

	if !compareCsvToStadium(lines[1], stadiums[0]) {
		t.FailNow()
	}

	if !compareCsvToStadium(lines[2], stadiums[1]) {
		t.FailNow()
	}

	ts.Close()
}

func TestGetStadiumsWithWeatherShowsUnavailableIfNotCached(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	ts := mockHttpServerWithRouter(client)
	// ingest our data
	err := ingestCsv(fmt.Sprintf("%s/api/stadiums", ts.URL), "", "")

	if err != nil {
		t.FailNow()
	}

	// Make a call to get our stadiums now

	resp, err := http.Get(fmt.Sprintf("%s/api/stadiums?weather=1", ts.URL))

	if err != nil {
		t.FailNow()
	}

	var decodedResponse stadiumResponse

	err = json.NewDecoder(resp.Body).Decode(&decodedResponse)

	if err != nil {
		t.Log("json decode failed")
		t.FailNow()
	}

	count := len(decodedResponse.Stadiums)

	if count != 2 {
		t.Logf("Expected count of 2, got %d instead", count)
		t.FailNow()
	}

	stadiums := decodedResponse.Stadiums

	expected := weather.Weather{
		Description: "unavailable",
		Icon:        "01d",
	}

	if *stadiums[0].Weather != expected {
		t.FailNow()
	}

	if *stadiums[1].Weather != expected {
		t.FailNow()
	}

	ts.Close()
}

func TestGetStadiumsWithWeatherShowsDataIfCached(t *testing.T) {
	// Set up a mock DB client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	ts := mockHttpServerWithRouter(client)
	// ingest our data
	err := ingestCsv(fmt.Sprintf("%s/api/stadiums", ts.URL), "", "")

	if err != nil {
		t.FailNow()
	}

	// insert some weather data
	ids := client.Stadium.Query().IDsX(context.Background())
	fmt.Println(ids)
	client.Weather.Create().SetIcon("13n").SetDescription("big snow storm").SetStadiumID(ids[0]).SetTemperature(-23).SaveX(context.Background())
	client.Weather.Create().SetIcon("10d").SetDescription("big rain storm").SetStadiumID(ids[1]).SetTemperature(15).SaveX(context.Background())

	// Make a call to get our stadiums now

	resp, err := http.Get(fmt.Sprintf("%s/api/stadiums?weather=1", ts.URL))

	if err != nil {
		t.FailNow()
	}

	var decodedResponse stadiumResponse

	err = json.NewDecoder(resp.Body).Decode(&decodedResponse)

	if err != nil {
		t.Log("json decode failed")
		t.FailNow()
	}

	count := len(decodedResponse.Stadiums)

	if count != 2 {
		t.Logf("Expected count of 2, got %d instead", count)
		t.FailNow()
	}

	stadiums := decodedResponse.Stadiums

	// sort by IDs to ensure the indexes are correct
	sort.Slice(stadiums, func(i, j int) bool {
		return stadiums[i].ID.String() < stadiums[j].ID.String()
	})

	if stadiums[0].Weather.Description != "big snow storm" {
		t.Logf("expected description of 'big snow storm' for stadium 1, got %s instead", stadiums[0].Weather.Description)
		t.FailNow()
	}

	if stadiums[0].Weather.Temp != -23 {
		t.Logf("expected temp of -23 for stadium 1, got %f instead", stadiums[0].Weather.Temp)
		t.FailNow()
	}

	if stadiums[0].Weather.Icon != "13n" {
		t.Logf("expected icon of '13n' for stadium 1, got %s instead", stadiums[0].Weather.Icon)
		t.FailNow()
	}

	if stadiums[1].Weather.Description != "big rain storm" {
		t.Logf("expected description of 'big rain storm' for stadium 2, got %s instead", stadiums[1].Weather.Description)
		t.FailNow()
	}

	if stadiums[1].Weather.Temp != 15 {
		t.Logf("expected temp of 15 for stadium 2, got %f instead", stadiums[1].Weather.Temp)
		t.FailNow()
	}

	if stadiums[1].Weather.Icon != "10d" {
		t.Logf("expected icon of '10d' for stadium 1, got %s instead", stadiums[1].Weather.Icon)
		t.FailNow()
	}

	ts.Close()
}
