package xphp

import "time"

// Checkdate checkdate()
// Validate a Gregorian date
func Checkdate(month, day, year int) bool {
	// check month
	if month > 12 || month < 1 {
		return false
	}
	// check day
	if day < 1 {
		return false
	}
	switch month {
	case 1, 3, 5, 7, 8, 10, 12: // 31 days
		if day > 31 {
			return false
		}
	case 4, 6, 9, 11: // 30 days
		if day > 30 {
			return false
		}
	default: // february
		if CheckIfLeapYear(year) {
			if day > 29 {
				return false
			}
		} else {
			if day > 28 {
				return false
			}
		}
	}
	// check year
	if year < 1 || year > 32767 {
		return false
	}
	return true
}

func CheckIfLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	if year%4 == 0 {
		return true
	}
	return false
}

// Time time()
func Time() int64 {
	return time.Now().Unix()
}

// Strtotime strtotime()
// Strtotime("02/01/2006 15:04:05", "02/01/2016 15:04:05") == 1451747045
// Strtotime("3 04 PM", "8 41 PM") == -62167144740
func Strtotime(format, strtime string) (int64, error) {
	t, err := time.Parse(format, strtime)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func StrToTime(str string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// Date date()
// Date("02/01/2006 15:04:05 PM", 1524799394)
func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}

// Sleep sleep()
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// Usleep usleep()
func Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
