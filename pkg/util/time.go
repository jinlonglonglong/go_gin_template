package util

import (
	"fmt"
	"time"
)

func GetTodayString() string {
	return time.Now().Format("20060102")
}

func GetDayOfMonthString() string {
	return fmt.Sprintf("%d", time.Now().Day())
}

func GetTimestampOfAMonthAgo() int64 {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	//hour, min, sec := now.UTC().Clock()
	newTime := time.Date(year, mon, day, 0, 0, 0, 0, time.UTC)
	return newTime.AddDate(0, -1, -1).Unix()
}

func GetZeroTimestampOfToday() int64 {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	//hour, min, sec := now.UTC().Clock()
	newTime := time.Date(year, mon, day, 0, 0, 0, 0, time.UTC)

	return newTime.Unix()
}

func TransferTimestampToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).UTC().Format("2006-01-02")
}
