package carbon

import (
	"strconv"
	"strings"
)

// Parse 将标准格式时间字符串解析成 Carbon 实例
func (c Carbon) Parse(value string) Carbon {
	if c.Error != nil {
		return c
	}

	layout := DateTimeFormat

	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}

	if len(value) == 10 && strings.Count(value, "-") == 2 {
		layout = DateFormat
	}

	if strings.Index(value, "T") == 10 {
		layout = RFC3339Format
	}

	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		switch len(value) {
		case 8:
			layout = ShortDateFormat
		case 14:
			layout = ShortDateTimeFormat
		}
	}

	tt, err := parseByLayout(value, layout)
	c.Time = tt
	c.Error = err
	return c
}

// Parse 将标准格式时间字符串解析成 Carbon 实例(默认时区)
func Parse(value string) Carbon {
	return SetTimezone(Local).Parse(value)
}

// ParseByFormat 将特殊格式时间字符串解析成 Carbon 实例
func (c Carbon) ParseByFormat(value string, format string) Carbon {
	if c.Error != nil {
		return c
	}
	layout := format2layout(format)
	return c.ParseByLayout(value, layout)
}

// ParseByFormat 将特殊格式时间字符串解析成 Carbon 实例(默认时区)
func ParseByFormat(value string, format string) Carbon {
	return SetTimezone(Local).ParseByFormat(value, format)
}

// ParseByLayout 将布局时间字符串解析成 Carbon 实例
func (c Carbon) ParseByLayout(value string, layout string) Carbon {
	if c.Error != nil {
		return c
	}
	tt, err := parseByLayout(value, layout)
	c.Time = tt
	c.Error = err
	return c
}

// ParseByLayout 将布局时间字符串解析成 Carbon 实例(默认时区)
func ParseByLayout(value string, layout string) Carbon {
	return SetTimezone(Local).ParseByLayout(value, layout)
}
