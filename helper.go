package carbon

import (
	"bytes"
	"errors"
	"time"
)

// 常规格式化符号
var formats = map[byte]string{
	'd': "02",                        // Day:    Day of the month, 2 digits with leading zeros. Eg: 01 to 31.
	'D': "Mon",                       // Day:    A textual representation of a day, three letters. Eg: Mon through Sun.
	'j': "2",                         // Day:    Day of the month without leading zeros. Eg: 1 to 31.
	'l': "Monday",                    // Day:    A full textual representation of the day of the week. Eg: Sunday through Saturday.
	'F': "January",                   // Month:  A full textual representation of a month, such as January or March. Eg: January through December.
	'm': "01",                        // Month:  Numeric representation of a month, with leading zeros. Eg: 01 through 12.
	'M': "Jan",                       // Month:  A short textual representation of a month, three letters. Eg: Jan through Dec.
	'n': "1",                         // Month:  Numeric representation of a month, without leading zeros. Eg: 1 through 12.
	'Y': "2006",                      // Year:   A full numeric representation of a year, 4 digits. Eg: 1999 or 2003.
	'y': "06",                        // Year:   A two digit representation of a year. Eg: 99 or 03.
	'a': "pm",                        // Time:   Lowercase Ante meridiem and Post meridiem. Eg: am or pm.
	'A': "PM",                        // Time:   Uppercase Ante meridiem and Post meridiem. Eg: AM or PM.
	'g': "3",                         // Time:   12-hour format of an hour without leading zeros. Eg: 1 through 12.
	'h': "03",                        // Time:   12-hour format of an hour with leading zeros. Eg: 01 through 12.
	'H': "15",                        // Time:   24-hour format of an hour with leading zeros. Eg: 00 through 23.
	'i': "04",                        // Time:   Minutes with leading zeros. Eg: 00 to 59.
	's': "05",                        // Time:   Seconds with leading zeros. Eg: 00 through 59.
	'O': "-0700",                     // Zone:   Difference to Greenwich time (GMT) in hours. Eg: +0200.
	'P': "-07:00",                    // Zone:   Difference to Greenwich time (GMT) with colon between hours and minutes. Eg: +02:00.
	'T': "MST",                       // Zone:   Timezone abbreviation. Eg: UTC, EST, MDT ...
	'c': "2006-01-02T15:04:05-07:00", // Format: ISO 8601 date. Eg: 2004-02-12T15:19:21+00:00.
	'r': "Mon, 02 Jan 06 15:04 MST",  // Format: RFC 2822 formatted date. Eg: Thu, 21 Dec 2000 16:01:07 +0200.
}

// format转layout
func format2layout(format string) string {
	runes := []rune(format)
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len(runes); i++ {
		if layout, ok := formats[byte(runes[i])]; ok {
			buffer.WriteString(layout)
		} else {
			buffer.WriteRune(runes[i])
		}
	}
	return buffer.String()
}

// getLocationByTimezone 通过时区获取Location实例
func getLocationByTimezone(timezone string) (*time.Location, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		err = errors.New("invalid timezone \"" + timezone + "\", please see the $GOROOT/lib/time/zoneinfo.zip file for all valid timezone")
	}
	return loc, err
}

// parseByLayout 通过布局模板解析
func parseByLayout(value string, layout string) (time.Time, error) {
	loc, _ := time.LoadLocation(Local)
	tt, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		err = errors.New("the value \"" + value + "\" can't parse string as time")
	}
	return tt, err
}

// parseByDuration 通过持续时间解析
func parseByDuration(duration string) (time.Duration, error) {
	td, err := time.ParseDuration(duration)
	if err != nil {
		err = errors.New("invalid duration \"" + duration + "\"")
	}
	return td, err
}

// getAbsValue 获取绝对值
func getAbsValue(value int64) int64 {
	return (value ^ value>>31) - value>>31
}
