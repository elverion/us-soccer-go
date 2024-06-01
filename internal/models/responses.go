package models

type Location struct {
	City    string  `json:"city"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
}

type Weather struct {
	Temp        float64 `json:"temp"`
	Description string  `json:"description"`
}

type Stadium struct {
	Stadium  string   `json:"stadium"`
	Location Location `json:"location"`
	Weather  Weather  `json:"weather"`
}
