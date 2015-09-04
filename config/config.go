package config
import (
	"os"
	"log"
	"net/url"
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

func GetRedisHostAndPassword() (string, string) {
	_url, err := url.Parse(Get("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}
	host := _url.Host
	finalPwd := ""
	user := _url.User
	if user == nil {
		log.Println("Redis user info not set, will try to use the empty string as the password")
	} else {
		pwd, pwdSet := _url.User.Password()
		if !pwdSet {
			log.Println("Redis password not set, will try to use the empty string")
		} else {
			finalPwd = pwd
		}
	}

	log.Printf("Redis host : %s, Redis password : %s", host, finalPwd)
	return host, finalPwd
}

