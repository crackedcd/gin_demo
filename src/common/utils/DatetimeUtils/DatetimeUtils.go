package DatetimeUtils

import "time"

var formatStr = "2006-01-02 15:04:05"

func GetTime() time.Time {
	return time.Now()
}

// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func GetTimeDuration(s string) time.Time {
	now := time.Now()
	d, _ := time.ParseDuration(s)
	return now.Add(d)
}

func GetTimeStrDuration(s string) string {
	return GetTimeDuration(s).Format(formatStr)
}

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func GetTimestampNano() int64 {
	return time.Now().UnixNano()
}

func GetDatetime() string {
	return time.Now().Format(formatStr)
}

func TimestampToStr(ts int64) string {
	return time.Unix(ts, 0).Format(formatStr)
}

func StrToTimestamp(s string) int64 {
	return StrToTime(s).Unix()
}

func StrToTime(s string) time.Time {
	t, _ := time.Parse(formatStr, s)
	return t
}
