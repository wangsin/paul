package sysLogic

import timeUtil "github.com/wangsin/paul/util/time"

func GetTime() map[string]interface{} {
	utc := timeUtil.GetNowUTCTime()
	local := timeUtil.GetNowLocalTime()
	return map[string]interface{}{
		"utcTime":            timeUtil.FormatWith(utc, timeUtil.TimeOnYMDHMS),
		"utcTimestamp":       timeUtil.ToTimeStamp(utc, false),
		"utcNanoTimestamp":   timeUtil.ToTimeStamp(utc, true),
		"localTime":          timeUtil.FormatWith(local, timeUtil.TimeOnYMDHMS),
		"localTimestamp":     timeUtil.ToTimeStamp(local, false),
		"localNanoTimestamp": timeUtil.ToTimeStamp(local, true),
	}
}
