package utils

import (
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

func EndOfLastMonthJalaali() time.Time {
	return ptime.Now().BeginningOfMonth().Add(-1 * time.Second).Time().In(time.UTC)
}
