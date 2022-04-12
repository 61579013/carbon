package carbon

import (
	"time"
)

// CreateFromTimestampWithSecond creates a Carbon instance from a given timestamp with second.
// 从给定的秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromTimestampWithSecond(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp, 0)
	return c
}

// CreateFromTimestampWithSecond creates a Carbon instance from a given timestamp with second.
// 从给定的秒级时间戳创建 Carbon 实例
func CreateFromTimestampWithSecond(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestampWithSecond(timestamp, timezone...)
}

// CreateFromTimestamp creates a Carbon instance from a given timestamp with second, it is shorthand for CreateFromTimestampWithSecond.
// 从给定的秒级时间戳创建 Carbon 实例, 是 CreateFromTimestampWithSecond 的简写
func (c Carbon) CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	return c.CreateFromTimestampWithSecond(timestamp, timezone...)
}

// CreateFromTimestamp creates a Carbon instance from a given timestamp with second, it is shorthand for CreateFromTimestampWithSecond.
// 从给定的秒级时间戳创建 Carbon 实例, 是 CreateFromTimestampWithSecond 的简写
func CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	return CreateFromTimestampWithSecond(timestamp, timezone...)
}

// CreateFromTimestampWithMillisecond creates a Carbon instance from a given timestamp with millisecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromTimestampWithMillisecond(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp/1e3, (timestamp%1e3)*1e6)
	return c
}

// CreateFromTimestampWithMillisecond creates a Carbon instance from a given timestamp with millisecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func CreateFromTimestampWithMillisecond(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestampWithMillisecond(timestamp, timezone...)
}

// CreateFromTimestampMilli creates a Carbon instance from a given timestamp with millisecond, it is shorthand for CreateFromTimestampWithMillisecond.
// 从给定的微秒时间戳创建 Carbon 实例, 是 CreateFromTimestampWithMillisecond 的简写
func (c Carbon) CreateFromTimestampMilli(timestamp int64, timezone ...string) Carbon {
	return c.CreateFromTimestampWithMillisecond(timestamp, timezone...)
}

// CreateFromTimestampMilli creates a Carbon instance from a given timestamp with millisecond, it is shorthand for CreateFromTimestampWithMillisecond.
// 从给定的微秒时间戳创建 Carbon 实例, 是 CreateFromTimestampWithMillisecond 的简写
func CreateFromTimestampMilli(timestamp int64, timezone ...string) Carbon {
	return CreateFromTimestampWithMillisecond(timestamp, timezone...)
}

// CreateFromTimestampWithMicrosecond creates a Carbon instance from a given timestamp with microsecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromTimestampWithMicrosecond(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp/1e6, (timestamp%1e6)*1e3)
	return c
}

// CreateFromTimestampWithMicrosecond creates a Carbon instance from a given timestamp with microsecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func CreateFromTimestampWithMicrosecond(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestampWithMicrosecond(timestamp, timezone...)
}

// CreateFromTimestampMicro creates a Carbon instance from a given timestamp with microsecond, it is shorthand for CreateFromTimestampWithMicrosecond.
// 从给定的微秒时间戳创建 Carbon 实例, 是 CreateFromTimestampMicro 的简写
func (c Carbon) CreateFromTimestampMicro(timestamp int64, timezone ...string) Carbon {
	return c.CreateFromTimestampWithMicrosecond(timestamp, timezone...)
}

// CreateFromTimestampMicro creates a Carbon instance from a given timestamp with microsecond, it is shorthand for CreateFromTimestampWithMicrosecond.
// 从给定的微秒时间戳创建 Carbon 实例, 是 CreateFromTimestampMicro 的简写
func CreateFromTimestampMicro(timestamp int64, timezone ...string) Carbon {
	return CreateFromTimestampWithMicrosecond(timestamp, timezone...)
}

// CreateFromTimestampWithNanosecond creates a Carbon instance from a given timestamp with nanosecond.
// 从给定的纳秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromTimestampWithNanosecond(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp/1e9, timestamp%1e9)
	return c
}

// CreateFromTimestampWithNanosecond creates a Carbon instance from a given timestamp with nanosecond.
// 从给定的纳秒级时间戳创建 Carbon 实例
func CreateFromTimestampWithNanosecond(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestampWithNanosecond(timestamp, timezone...)
}

// CreateFromTimestampNano creates a Carbon instance from a given timestamp with nanosecond, it is shorthand for CreateFromTimestampWithNanosecond.
// 从给定的纳秒时间戳创建 Carbon 实例, 是 CreateFromTimestampWithNanosecond 的简写
func (c Carbon) CreateFromTimestampNano(timestamp int64, timezone ...string) Carbon {
	return c.CreateFromTimestampWithNanosecond(timestamp, timezone...)
}

// CreateFromTimestampNano creates a Carbon instance from a given timestamp with nanosecond, it is shorthand for CreateFromTimestampWithNanosecond.
// 从给定的纳秒时间戳创建 Carbon 实例, 是 CreateFromTimestampWithNanosecond 的简写
func CreateFromTimestampNano(timestamp int64, timezone ...string) Carbon {
	return CreateFromTimestampWithNanosecond(timestamp, timezone...)
}

// Create creates a Carbon instance from a given time.
// 从给定的时间创建 Carbon 实例
func (c Carbon) Create(year, month, day, hour, minute, second, nanosecond int, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, c.loc)
	return c
}

// Create creates a Carbon instance from a given time.
// 从给定的时间创建 Carbon 实例
func Create(year, month, day, hour, minute, second, nanosecond int, timezone ...string) Carbon {
	return NewCarbon().Create(year, month, day, hour, minute, second, nanosecond, timezone...)
}

// CreateFromDateTime creates a Carbon instance from a given date and time.
// 从给定的年月日时分秒创建 Carbon 实例
func (c Carbon) CreateFromDateTime(year, month, day, hour, minute, second int, timezone ...string) Carbon {
	return c.Create(year, month, day, hour, minute, second, time.Now().Nanosecond(), timezone...)
}

// CreateFromDateTime creates a Carbon instance from a given date and time.
// 从给定的年月日时分秒创建 Carbon 实例
func CreateFromDateTime(year, month, day, hour, minute, second int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDateTime(year, month, day, hour, minute, second, timezone...)
}

// CreateFromDate creates a Carbon instance from a given date.
// 从给定的年月日创建 Carbon 实例
func (c Carbon) CreateFromDate(year, month, day int, timezone ...string) Carbon {
	now := Now()
	hour, minute, second := now.Time()
	return c.Create(year, month, day, hour, minute, second, time.Now().Nanosecond(), timezone...)
}

// CreateFromDate creates a Carbon instance from a given date.
// 从给定的年月日创建 Carbon 实例
func CreateFromDate(year, month, day int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDate(year, month, day, timezone...)
}

// CreateFromTime creates a Carbon instance from a given time.
// 从给定的时分秒创建 Carbon 实例
func (c Carbon) CreateFromTime(hour, minute, second int, timezone ...string) Carbon {
	now := Now()
	year, month, day := now.Date()
	return c.Create(year, month, day, hour, minute, second, now.Nanosecond(), timezone...)
}

// CreateFromTime creates a Carbon instance from a given time.
// 从给定的时分秒创建 Carbon 实例
func CreateFromTime(hour, minute, second int, timezone ...string) Carbon {
	return NewCarbon().CreateFromTime(hour, minute, second, timezone...)
}
