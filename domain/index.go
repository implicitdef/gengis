package domain
import (
	"fmt"
	"time"
)



type TimeRange struct {
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
}

func (tr TimeRange) String() string {
	return fmt.Sprintf("{%s, %s}", tr.Start.String(), tr.End.String())
}