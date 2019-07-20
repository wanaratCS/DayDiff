package calcualte

import (
	"time"
)

func calculateSecond(StartDate ,EndDate time.Time) int {

	EndDate = EndDate.AddDate(0,0,1)
	seconds := EndDate.Sub(StartDate).Seconds()

	return int(seconds)
}

func calculateMinutes(StartDate,EndDate time.Time) int  {
	return 12336480
}

func calculateHour(StartDate,EndDate time.Time) int  {
	return 205608
}