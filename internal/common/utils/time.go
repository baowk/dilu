package utils

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

func GetZoreTimeLocal(t time.Time) time.Time {
	t = t.Local()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func GetWeekFirstDayLocal(t time.Time) time.Time {
	t = t.Local()
	return time.Date(t.Year(), t.Month(), t.Day()-int(t.Weekday()), 0, 0, 0, 0, time.Local)
}

func GetMonthFirstDayLocal(t time.Time) time.Time {
	t = t.Local()
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

// func GetZoreTime(t time.Time) time.Time {
// 	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
// }

// func GetWeekFirstDay(t time.Time) time.Time {
// 	return time.Date(t.Year(), t.Month(), t.Day()-int(t.Weekday()), 0, 0, 0, 0, time.Local)
// }

// func GetMonthFirstDay(t time.Time) time.Time {
// 	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
// }

func GetZoreTimeLocation(t time.Time, location *time.Location) time.Time {
	t = t.In(location)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func GetWeekFirstDayLocation(t time.Time, location *time.Location) time.Time {
	t = t.In(location)
	return time.Date(t.Year(), t.Month(), t.Day()-int(t.Weekday()), 0, 0, 0, 0, time.Local)
}

func GetMonthFirstDayLocation(t time.Time, location *time.Location) time.Time {
	t = t.In(location)
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

var MonthDay = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func GetMonthLen(t time.Time) int {
	t = t.Local()
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

func CmpAge(birthday time.Time) int {
	return time.Now().Year() - birthday.Year()
}

type LocalTime time.Time

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *LocalTime) String() string {
	// 如果时间 null 那么我们需要把返回的值进行修改
	if t == nil || t.IsZero() {
		return ""
	}
	return fmt.Sprintf("%s", time.Time(*t).Format("2006-01-02 15:04:05"))
}

func (t *LocalTime) IsZero() bool {
	return time.Time(*t).IsZero()
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	*t = LocalTime(t1)
	return err
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t)
	// 如果时间值是空或者0值 返回为null 如果写空字符串会报错
	if &t == nil || t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", tTime.Format("2006-01-02 15:04:05"))), nil
}
