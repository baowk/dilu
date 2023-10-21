package utils

import "time"

func GetZoreTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func GetWeekFirstDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()-int(t.Weekday()), 0, 0, 0, 0, time.Local)
}

func GetMonthFirstDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

var MonthDay = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func GetMonthLen(t time.Time) int {
	month := int(t.Month())
	if month == 2 {
		year := t.Year()
		// 计算是平年还是闰年
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			// 得出2月的天数
			return 29
		}
		// 得出2月的天数
		return 28
	}
	return MonthDay[month-1]
}
