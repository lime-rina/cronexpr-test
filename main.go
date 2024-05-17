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
	// CronHelper("* * 16 L-2 2/3 ? *", time.Date(3023, time.August, 30, 0, 0, 0, 0, time.UTC))
	// n := CronHelper("0 0 0 L-3 * * ?", time.Now())
	n := CronHelper("0 0 29 2 *", time.Date(2323, time.January, 30, 0, 0, 0, 0, time.UTC))
	// from, _ := time.Parse("2006-01-02", "2013-08-31")
	// fmt.Println(from)
	// n := cronexpr.MustParse("* * * * * 1980").Next(from)

	CronHelper2("0 0 0 * * THUL ?", time.Date(2124, time.August, 30, 0, 0, 0, 0, time.UTC))
	fmt.Println("next: ", n)

}

func CronHelper(cronExpression string, startTime time.Time) time.Time {
	// nextTime := cronexpr.MustParse(cronExpression).NextN(startTime, 3)
	nextTime := cronexpr.MustParse(cronExpression).Next(startTime)
	// nextTime := cronexpr.MustParse(cronExpression).NextN(startTime, 5)
	// // fmt.Println("must parse: ", cronexpr.MustParse(cronExpression))

	// for i := range nextTime {
	// 	fmt.Println(nextTime[i])
	// }
	return nextTime
}

func CronHelper2(cronExpression string, startTime time.Time) []time.Time {
	// nextTime := cronexpr.MustParse(cronExpression).NextN(startTime, 3)
	nextTime := cronexpr.MustParse(cronExpression).NextN(startTime, 4)
	// nextTime := cronexpr.MustParse(cronExpression).NextN(startTime, 5)
	// fmt.Println("must parse: ", cronexpr.MustParse(cronExpression))

	for i := range nextTime {
		fmt.Println(nextTime[i])
	}
	return nextTime
}
