package tibber

import "os"

func GetMyApiKey() string {
	key := os.Getenv("TIBBER_API_KEY")

	if key == "" {
		panic("TIBBER_API_KEY is not set")
	}
	return key
}
