package main

import (
	"fmt"
	"time"

	"github.com/hashicorp/cronexpr"
)

func main() {
	// CronHelper("0 2 1 * * *", time.Date(2022, time.November, 1, 1, 1, 0, 0, time.UTC))
	// CronHelper("0 0 12 * *", time.Date(2122, time.November, 1, 1, 1, 0, 0, time.UTC))
	// CronHelper("0 0 17 LW 1/3 ? *", time.Now())
	// CronHelper("* * 16 L-2 2/3 ? *", time.Date(2024, time.August, 30, 0, 0, 0, 0, time.UTC))
	// CronHelper("0 0 0 L-3 * * ?", time.Now())

	CronHelper("0 0 0 * * THUL ?", time.Now())

}

func CronHelper(cronExpression string, startTime time.Time) time.Time {
	// nextTime := cronexpr.MustParse(cronExpression).NextN(startTime, 3)
	nextTime := cronexpr.MustParse(cronExpression).Next(startTime)

	fmt.Println("nextTime: ", nextTime)
	return nextTime
}
