package time_t

import (
	"strings"
	"time"
)

const DateTimeFormat = "2006-01-02T15:04:05.999999999Z070"

func convertStrToTime(target string) time.Time {
	loc, _ := time.LoadLocation("Europe/Moscow")
	t, _ := time.ParseInLocation(DateTimeFormat, target, loc)
	return t
}

func convertTimeToStr(target time.Time) string {
	timeStr := target.Format(DateTimeFormat)
	if target.Nanosecond()-1e6 == 0 {
		timeStr = strings.Replace(timeStr, "+", ".000+", -1)
	}
	return timeStr
}
