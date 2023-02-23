package config

import "os"

var BaseURL = os.Getenv("BASE_URL")

func SetURL(url string) string {
	if BaseURL == "" {
		BaseURL = "http://localhost:3000"
	}

	return BaseURL + url
}
