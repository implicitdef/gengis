package rediscache
import (
	"gopkg.in/redis.v3"
	"encoding/json"
	"time"
	"github.com/mtailor/gengis/myrandom"
	"github.com/mtailor/gengis/myerrors"
	"github.com/mtailor/gengis/config"
)

const prefix = "cache:"

var host, pwd = config.GetRedisHostAndPassword()

var client = redis.NewClient(&redis.Options{
	Addr: host,
	Password: pwd,
	DB : 0,
})

// This packages assumes the given values are perfectly
// marshable/unmarshable in JSON (so no unexported fields...)

// marshalles the value and puts it in this key
// will expire automatically after an arbitrary time
func Set(key string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	duration := time.Duration((myrandom.Generator.Intn(10) + 10)) * time.Hour
	return client.Set(prefix + key, bytes, duration).Err()
}

// returns the value that was marshalled into this key
// error may be NotInCacheError if it's not there (not set or expired),
// or may also be any unpredictable kind of error
func Get(key string, dest interface{}) error {
	res := client.Get(prefix + key)
	bytes, err := res.Bytes()
	if err == redis.Nil {
		return &myerrors.NotInCacheError{"The key " + key + " was not cached"}
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dest)
}
