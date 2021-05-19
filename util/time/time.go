package timeUtil

import "time"

type TimeFormat string

const (
	TimeOnYMDHMS TimeFormat = "2006-01-02 15:04:05"
)

func GetNowUTCTime() time.Time {
	return time.Now().UTC()
}

func GetNowLocalTime() time.Time {
	return time.Now().Local()
}

func ToTimeStamp(t time.Time, nano bool) int64 {
	if !nano {
		return t.Unix()
	}

	return t.UnixNano()
}

func FormatWithRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func FormatWith(t time.Time, s TimeFormat) string {
	return t.Format(string(s))
}
