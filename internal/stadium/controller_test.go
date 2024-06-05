package stadium_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"us-soccer-go-test/internal/ent"
	"us-soccer-go-test/internal/ent/enttest"
	"us-soccer-go-test/internal/stadium"

	"github.com/apex/log"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

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
