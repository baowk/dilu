package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestGetZoreTime(t *testing.T) {
	time.Parse(time.RFC1123, "2023-10-29 16:00:00 +0000 UTC")
	cur := time.Now()

	fmt.Println(GetZoreTimeLocal(cur))

	cur = cur.Add(24 * time.Hour)
	fmt.Println(GetZoreTimeLocal(cur))

	cur = cur.Add(-24 * time.Hour)
	fmt.Println(GetZoreTimeLocal(cur))
	cur = cur.Add(-24 * 7 * time.Hour)
	fmt.Println(GetZoreTimeLocal(cur))

	fmt.Println(time.Sunday)
}

func TestGetWeekFirstDay(t *testing.T) {
	cur := time.Now()
	fmt.Println(cur)
	fmt.Println(GetWeekFirstDayLocal(cur))

	cur = cur.Add(-24 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetWeekFirstDayLocal(cur))

	cur = cur.Add(-24 * 7 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetWeekFirstDayLocal(cur))

	cur = cur.Add(-24 * 12 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetWeekFirstDayLocal(cur))
}

func TestGetMonthFirstDay(t *testing.T) {
	cur := time.Now()
	fmt.Println(cur)
	fmt.Println(GetMonthFirstDayLocal(cur))

	cur = cur.Add(-24 * 7 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetMonthFirstDayLocal(cur))

	cur = cur.Add(-24 * 7 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetMonthFirstDayLocal(cur))

	cur = cur.Add(-24 * 7 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetMonthFirstDayLocal(cur))
}

func TestGetMonthLen(t *testing.T) {
	cur := time.Now()
	fmt.Println(cur)
	fmt.Println(GetMonthLen(cur))

	cur = cur.Add(-24 * 231 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetMonthLen(cur))

	cur = cur.Add(-24 * 31 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetMonthLen(cur))

	cur = cur.Add(-24 * 365 * 3 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetMonthLen(cur))

	cur = cur.Add(-24 * 298 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetMonthLen(cur))

	cur = cur.Add(-24 * 298 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetMonthLen(cur))
}
