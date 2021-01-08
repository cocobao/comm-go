package utility

import "time"

func ExTimeSimple(t time.Time) string {
	return t.Local().Format("2006-01-02 15:04:05")
}

func ExTime(t time.Time) string {
	return t.Format("2006-01-02T15:04:05-07:00")
}

func ExLocalTime(t time.Time) string {
	return t.Local().Format("2006-01-02T15:04:05-07:00")
}

func ExTimeDetail(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.999999-07:00")
}

func ExTimeDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func NowTDetail() string {
	return time.Now().Format("2006-01-02T15:04:05.999999-07:00")
}

func StringToTimeDetail(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02T15:04:05.999999-07:00", s, loc)
	return theTime
}

func StringToTime(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02T15:04:05-07:00", s, loc)
	return theTime
}

func StringToDate(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02", s, loc)
	return theTime
}

func NowT() string {
	return time.Now().Format("2006-01-02T15:04:05-07:00")
}

func NowDate() string {
	return time.Now().Format("0102")
}

func NowYearDate() string {
	return time.Now().Format("20060102")
}

func NowTime() string {
	return time.Now().Format("20060102150405")
}

func NowN() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func NowDateName() string {
	return time.Now().Format("2006/01/02")
}

func NowDateTime() string {
	return time.Now().Format("01-02 15:04:05")
}

func GetTimeZone() string {
	return time.Now().Format("-0700")
}

func TimestampToString(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02T15:04:05-07:00")
}

func TimeIntToTimeStamp(t int64) time.Time {
	sec := t / 1000
	ns := t % 1000
	ns *= 1000000
	return time.Unix(sec, ns)
}

func TimeIntToString(t int64) string {
	sec := t / 1000
	ns := t % 1000
	ns *= 1000000
	return time.Unix(sec, ns).Format("2006-01-02T15:04:05-07:00")
}

func TimeIntSecToTimeStamp(t int64) time.Time {
	return time.Unix(t, 0)
}
