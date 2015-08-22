package domain
import (
	"fmt"
	"time"
)



type TimeRange struct {
	Start time.Time
	End time.Time
}

func (tr TimeRange) String() string {
	return fmt.Sprintf("{%s, %s}", tr.Start.String(), tr.End.String())
}