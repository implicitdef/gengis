package config
import (
	"os"
	"log"
)

func getWithFallBack(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}

func GetTheMovieDbApiKey() string {
	return getWithFallBack("THEMOVIEDB_API_KEY", "000ff" + "c8b6e767158" + "ff5489a8daba11c2")
}

func Get(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s must be set", key)
	}
	log.Printf("Using %s = %s", key, value)
	return value
}