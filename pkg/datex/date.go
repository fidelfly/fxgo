package system

import (
	"time"
)

const (
	dateFormat = "2006-01-02"
	//TimeFormat = "2006-01-02 15:04:05"
)

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+dateFormat+`"`, string(data), time.Local)
	*d = Date(now)
	return
}

func (d Date) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		return []byte("\"\""), nil
	}
	b := make([]byte, 0, len(dateFormat)+2)
	b = append(b, '"')
	b = time.Time(d).AppendFormat(b, dateFormat)
	b = append(b, '"')
	return b, nil
}

func Today() time.Time {
	return DateStart(time.Now())
}

func TodayEnd() time.Time {
	return DateEnd(time.Now())
}

func DateStart(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func DateEnd(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999, date.Location())
}
