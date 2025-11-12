package db

import (
	"fmt"
	"time"

	"database/sql/driver"
)

// CustomTime 自定义时间类型
type CustomTime time.Time

func (t CustomTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("\"\""), nil
	}
	b := make([]byte, 0, len(time.DateTime)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, time.DateTime)
	b = append(b, '"')
	return b, nil
}

func (t *CustomTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) > 2 && data[0] == '"' {
		now, err := time.ParseInLocation(`"`+time.DateTime+`"`, string(data), time.Local)
		*t = CustomTime(now)
		return err
	}
	return
}

func (t CustomTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tTime := time.Time(t)
	// 判断给定时间是否和默认零时间的时间戳相同
	if tTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tTime, nil
}

func (t *CustomTime) Scan(v any) error {
	value, ok := v.(time.Time)
	if ok {
		*t = CustomTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t CustomTime) String() string {
	if time.Time(t).IsZero() {
		return ""
	}
	return time.Time(t).Format(time.DateTime)
}

func (t CustomTime) ToTime() time.Time {
	tTime := time.Time(t)
	if tTime.IsZero() {
		return tTime
	}
	year, month, day := tTime.Date()
	return time.Date(year, month, day, tTime.Hour(), tTime.Minute(), tTime.Second(), 0, time.Local)
}

// CustomDate 自定义日期类型
type CustomDate time.Time

func (t CustomDate) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("\"\""), nil
	}
	b := make([]byte, 0, len(time.DateOnly)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, time.DateOnly)
	b = append(b, '"')
	return b, nil
}

func (t *CustomDate) UnmarshalJSON(data []byte) (err error) {
	if len(data) > 2 && data[0] == '"' {
		now, err := time.ParseInLocation(`"`+time.DateOnly+`"`, string(data), time.Local)
		*t = CustomDate(now)
		return err
	}
	return
}

func (t CustomDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	tTime := time.Time(t)
	// 判断给定时间是否和默认零时间的时间戳相同
	if tTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tTime, nil
}

func (t *CustomDate) Scan(v any) error {
	if value, ok := v.(time.Time); ok {
		*t = CustomDate(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t CustomDate) String() string {
	if time.Time(t).IsZero() {
		return ""
	}
	return time.Time(t).Format(time.DateOnly)
}

func (t CustomDate) ToTime() time.Time {
	tTime := time.Time(t)
	if tTime.IsZero() {
		return tTime
	}
	year, month, day := tTime.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}
