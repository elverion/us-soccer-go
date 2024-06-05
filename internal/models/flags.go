package models

type Flags struct {
	Env            string     `short:"e" long:"env" env:"ENV" required:"true" description:"The environment the program is running in: production/development"`
	HTTP           ConfigHTTP `group:"HTTP Server Options" namespace:"http" env-namespace:"HTTP"`
	APIKeys        APIKeys    `group:"API keys" namespace:"api-keys" env-namespace:"API_KEYS"`
	DB             ConfigDB   `group:"Database Options" namespace:"db" env-namespace:"DB"`
	RunDaemon      bool       `short:"d" long:"daemon" env:"RUN_DAEMON" description:"Determines if the weather daemon should be run"`
	OpenWeatherURL string     `short:"o" long:"open-weather-url" env:"OPENWEATHER_URL" required:"true" description:"The URL for the openweather url"`
}

type ConfigHTTP struct {
	Addr           string   `short:"a" long:"addr" default:":8080" env:"ADDR" description:"ip:port pair to bind to" required:"true"`
	TrustedProxies []string `long:"trusted-proxies" env:"TRUSTED_PROXIES" description:"set of CIDR ranges that are allowed to provide an X-Forwarded-For header"`
}

type APIKeys struct {
	OpenWeather string `long:"open-weather" env:"OPENWEATHER" description:"OpenWeather API key" required:"true"`
}

type ConfigDB struct {
	URL string `env:"URL" long:"url" required:"true" description:"database connection url"`
}
