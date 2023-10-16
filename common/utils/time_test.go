package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestGetZoreTime(t *testing.T) {
	cur := time.Now()
	fmt.Println(GetZoreTime(cur))
	cur = cur.Add(-24 * time.Hour)
	fmt.Println(GetZoreTime(cur))
	cur = cur.Add(-24 * 7 * time.Hour)
	fmt.Println(GetZoreTime(cur))

	fmt.Println(time.Sunday)
}

func TestGetWeekFirstDay(t *testing.T) {
	cur := time.Now()
	fmt.Println(cur)
	fmt.Println(GetWeekFirstDay(cur))

	cur = cur.Add(-24 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetWeekFirstDay(cur))

	cur = cur.Add(-24 * 7 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetWeekFirstDay(cur))

	cur = cur.Add(-24 * 12 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetWeekFirstDay(cur))
}

func TestGetMonthFirstDay(t *testing.T) {
	cur := time.Now()
	fmt.Println(cur)
	fmt.Println(GetMonthFirstDay(cur))

	cur = cur.Add(-24 * 7 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetMonthFirstDay(cur))

	cur = cur.Add(-24 * 7 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetMonthFirstDay(cur))

	cur = cur.Add(-24 * 7 * time.Hour)
	fmt.Println(cur)
	fmt.Println(GetMonthFirstDay(cur))
}
