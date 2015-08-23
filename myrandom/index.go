package myrandom
import (
	"math/rand"
	"time"
)


var Generator = rand.New(rand.NewSource(time.Now().UnixNano()))