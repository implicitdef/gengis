package main
import (
	"github.com/mtailor/gengis/vendor/_nuts/gopkg.in/redis.v3"
	"fmt"
)


func main() {


	//TODO use the cache
	//TODO make redis conf configurable







	client := redis.NewClient(&redis.Options{
		Addr:	"localhost:6379",
		Password: "",
		DB : 0,
	})

	res, err := client.Ping().Result()
	fmt.Println(res, err)


	res = client.Set("key", "value", 0).Val()
	fmt.Println(res)

}

