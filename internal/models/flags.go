package models

type Flags struct {
	Configured bool       `long:"configured" env:"CONFIGURED" required:"true" description:"If set to false, the web application will exit, should be set to true when everything is configured correctly"`
	Env        string     `short:"e" long:"env" env:"ENV" required:"true" description:"The environment the program is running in: production/development"`
	HTTP       ConfigHTTP `group:"HTTP Server Options" namespace:"http" env-namespace:"HTTP"`
	APIKeys    APIKeys    `group:"API keys" namespace:"api-keys" env-namespace:"API_KEYS"`
}

type ConfigHTTP struct {
	Addr           string   `short:"a" long:"addr" default:":8080" env:"ADDR" description:"ip:port pair to bind to" required:"true"`
	BaseURL        string   `long:"base-url" required:"true" env:"BASE_URL"`
	TrustedProxies []string `long:"trusted-proxies" env:"TRUSTED_PROXIES" descriotion:"set of CIDR ranges that are allowed to provide an X-Forwarded-For header"`
}

type APIKeys struct {
	OpenWeather string `long:"open-weather" env:"OPENWEATHER" description:"OpenWeather API key" required:"true"`
}
