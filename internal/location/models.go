package location

type Location struct {
	City    string  `json:"city"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
}
