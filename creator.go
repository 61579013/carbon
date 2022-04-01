package carbon

import (
	"time"
)

// CreateFromTimestamp creates a Carbon instance from a given timestamp with second.
// 从给定的秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp, 0)
	return c
}

// CreateFromTimestamp creates a Carbon instance from a given timestamp with second.
// 从给定的秒级时间戳创建 Carbon 实例
func CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestamp(timestamp, timezone...)
}

// CreateFromMilliTimestamp creates a Carbon instance from a given timestamp with millisecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromMilliTimestamp(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp/1e3, (timestamp%1e3)*1e6)
	return c
}

// CreateFromMilliTimestamp creates a Carbon instance from a given timestamp with millisecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func CreateFromMilliTimestamp(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromMilliTimestamp(timestamp, timezone...)
}

// CreateFromMicroTimestamp creates a Carbon instance from a given timestamp with microsecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromMicroTimestamp(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp/1e6, (timestamp%1e6)*1e3)
	return c
}

// CreateFromMicroTimestamp creates a Carbon instance from a given timestamp with microsecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func CreateFromMicroTimestamp(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromMicroTimestamp(timestamp, timezone...)
}

// CreateFromNanoTimestamp creates a Carbon instance from a given timestamp with nanosecond.
// 从给定的纳秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromNanoTimestamp(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp/1e9, timestamp%1e9)
	return c
}

// CreateFromNanoTimestamp creates a Carbon instance from a given timestamp with nanosecond.
// 从给定的纳秒级时间戳创建 Carbon 实例
func CreateFromNanoTimestamp(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromNanoTimestamp(timestamp, timezone...)
}

// Create creates a Carbon instance from a given time.
// 从给定的时间创建 Carbon 实例
func (c Carbon) Create(year, month, day, hour, minute, second, Nanosecond int, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, Nanosecond, c.loc)
	return c
}

// Create creates a Carbon instance from a given time.
// 从给定的时间创建 Carbon 实例
func Create(year, month, day, hour, minute, second, Nanosecond int, timezone ...string) Carbon {
	return NewCarbon().Create(year, month, day, hour, minute, second, Nanosecond, timezone...)
}

// CreateFromDateTime creates a Carbon instance from a given date and time.
// 从给定的年月日时分秒创建 Carbon 实例
func (c Carbon) CreateFromDateTime(year, month, day, hour, minute, second int, timezone ...string) Carbon {
	return c.Create(year, month, day, hour, minute, second, time.Now().Nanosecond(), timezone...)
}

// CreateFromDateTime creates a Carbon instance from a given date and time.
// 从给定的年月日时分秒创建 Carbon 实例
func CreateFromDateTime(year int, month int, day int, hour int, minute int, second int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDateTime(year, month, day, hour, minute, second, timezone...)
}

// CreateFromDate creates a Carbon instance from a given date.
// 从给定的年月日创建 Carbon 实例
func (c Carbon) CreateFromDate(year int, month int, day int, timezone ...string) Carbon {
	hour, minute, second := time.Now().In(c.loc).Clock()
	return c.Create(year, month, day, hour, minute, second, time.Now().Nanosecond(), timezone...)
}

// CreateFromDate creates a Carbon instance from a given date.
// 从给定的年月日创建 Carbon 实例
func CreateFromDate(year int, month int, day int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDate(year, month, day, timezone...)
}

// CreateFromTime creates a Carbon instance from a given time.
// 从给定的时分秒创建 Carbon 实例
func (c Carbon) CreateFromTime(hour int, minute int, second int, timezone ...string) Carbon {
	year, month, day := time.Now().In(c.loc).Date()
	return c.Create(year, int(month), day, hour, minute, second, time.Now().Nanosecond(), timezone...)
}

// CreateFromTime creates a Carbon instance from a given time.
// 从给定的时分秒创建 Carbon 实例
func CreateFromTime(hour int, minute int, second int, timezone ...string) Carbon {
	return NewCarbon().CreateFromTime(hour, minute, second, timezone...)
}
