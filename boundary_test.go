package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_StartOfCentury(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-01 00:00:00", "2000-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2000-01-01 00:00:00"},
		{"2020-02-01 13:14:15", "2000-01-01 00:00:00"},
		{"2020-02-28", "2000-01-01 00:00:00"},
		{"2020-02-29", "2000-01-01 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).StartOfCentury()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_EndOfCentury(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-01 00:00:00", "2099-12-31 23:59:59"},
		{"2020-01-31 23:59:59", "2099-12-31 23:59:59"},
		{"2020-02-01 13:14:15", "2099-12-31 23:59:59"},
		{"2020-02-28", "2099-12-31 23:59:59"},
		{"2020-02-29", "2099-12-31 23:59:59"},
	}

	for index, test := range tests {
		c := Parse(test.input).EndOfCentury()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_StartOfDecade(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-31 23:59:59", "2020-01-01 00:00:00"},
		{"2021-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2029-01-31 23:59:59", "2020-01-01 00:00:00"},
		{"2020-02-01 13:14:15", "2020-01-01 00:00:00"},
		{"2020-02-28", "2020-01-01 00:00:00"},
		{"2020-02-29", "2020-01-01 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).StartOfDecade()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_EndOfDecade(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-31 23:59:59", "2029-12-31 23:59:59"},
		{"2021-01-01 00:00:00", "2029-12-31 23:59:59"},
		{"2029-01-31 23:59:59", "2029-12-31 23:59:59"},
		{"2020-02-01 13:14:15", "2029-12-31 23:59:59"},
		{"2020-02-28", "2029-12-31 23:59:59"},
		{"2020-02-29", "2029-12-31 23:59:59"},
	}

	for index, test := range tests {
		c := Parse(test.input).EndOfDecade()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_StartOfYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-01 00:00:00"},
		{"2020-02-01 13:14:15", "2020-01-01 00:00:00"},
		{"2020-02-28", "2020-01-01 00:00:00"},
		{"2020-02-29", "2020-01-01 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).StartOfYear()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_EndOfYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-01 00:00:00", "2020-12-31 23:59:59"},
		{"2020-01-31 23:59:59", "2020-12-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-12-31 23:59:59"},
		{"2020-02-28", "2020-12-31 23:59:59"},
		{"2020-02-29", "2020-12-31 23:59:59"},
	}

	for index, test := range tests {
		c := Parse(test.input).EndOfYear()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_StartOfQuarter(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2020-01-01 00:00:00"},
		{"2020-02-15 00:00:00", "2020-01-01 00:00:00"},
		{"2020-03-15 00:00:00", "2020-01-01 00:00:00"},
		{"2020-04-15 23:59:59", "2020-04-01 00:00:00"},
		{"2020-05-15 23:59:59", "2020-04-01 00:00:00"},
		{"2020-06-15 23:59:59", "2020-04-01 00:00:00"},
		{"2020-07-15 23:59:59", "2020-07-01 00:00:00"},
		{"2020-08-15 13:14:15", "2020-07-01 00:00:00"},
		{"2020-09-15 13:14:15", "2020-07-01 00:00:00"},
		{"2020-10-15", "2020-10-01 00:00:00"},
		{"2020-11-15", "2020-10-01 00:00:00"},
		{"2020-12-15", "2020-10-01 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).StartOfQuarter()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_EndOfQuarter(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2020-03-31 23:59:59"},
		{"2020-02-15 00:00:00", "2020-03-31 23:59:59"},
		{"2020-03-15 00:00:00", "2020-03-31 23:59:59"},
		{"2020-04-15 23:59:59", "2020-06-30 23:59:59"},
		{"2020-05-15 23:59:59", "2020-06-30 23:59:59"},
		{"2020-06-15 23:59:59", "2020-06-30 23:59:59"},
		{"2020-07-15 23:59:59", "2020-09-30 23:59:59"},
		{"2020-08-15 13:14:15", "2020-09-30 23:59:59"},
		{"2020-09-15 13:14:15", "2020-09-30 23:59:59"},
		{"2020-10-15", "2020-12-31 23:59:59"},
		{"2020-11-15", "2020-12-31 23:59:59"},
		{"2020-12-15", "2020-12-31 23:59:59"},
	}

	for index, test := range tests {
		c := Parse(test.input).EndOfQuarter()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_StartOfMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2020-01-01 00:00:00"},
		{"2020-02-15 00:00:00", "2020-02-01 00:00:00"},
		{"2020-03-15 00:00:00", "2020-03-01 00:00:00"},
		{"2020-04-15 23:59:59", "2020-04-01 00:00:00"},
		{"2020-05-15 23:59:59", "2020-05-01 00:00:00"},
		{"2020-06-15 23:59:59", "2020-06-01 00:00:00"},
		{"2020-07-15 23:59:59", "2020-07-01 00:00:00"},
		{"2020-08-15 13:14:15", "2020-08-01 00:00:00"},
		{"2020-09-15 13:14:15", "2020-09-01 00:00:00"},
		{"2020-10-15", "2020-10-01 00:00:00"},
		{"2020-11-15", "2020-11-01 00:00:00"},
		{"2020-12-15", "2020-12-01 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).StartOfMonth()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_EndOfMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2020-01-31 23:59:59"},
		{"2020-02-15 00:00:00", "2020-02-29 23:59:59"},
		{"2020-03-15 00:00:00", "2020-03-31 23:59:59"},
		{"2020-04-15 23:59:59", "2020-04-30 23:59:59"},
		{"2020-05-15 23:59:59", "2020-05-31 23:59:59"},
		{"2020-06-15 23:59:59", "2020-06-30 23:59:59"},
		{"2020-07-15 23:59:59", "2020-07-31 23:59:59"},
		{"2020-08-15 13:14:15", "2020-08-31 23:59:59"},
		{"2020-09-15 13:14:15", "2020-09-30 23:59:59"},
		{"2020-10-15", "2020-10-31 23:59:59"},
		{"2020-11-15", "2020-11-30 23:59:59"},
		{"2020-12-15", "2020-12-31 23:59:59"},
	}

	for index, test := range tests {
		c := Parse(test.input).EndOfMonth()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_StartOfWeek(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0000-00-00 00:00:00", ""},
		{"", ""},
		{"0000-00-00 00:00:00", ""},

		{"2021-06-13", "2021-06-13 00:00:00"},
		{"2021-06-14", "2021-06-13 00:00:00"},
		{"2021-06-18", "2021-06-13 00:00:00"},
		{"2021-06-19", "2021-06-13 00:00:00"},
		{"2021-06-20", "2021-06-20 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).StartOfWeek()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_EndOfWeek(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0000-00-00 00:00:00", ""},
		{"", ""},
		{"0000-00-00 00:00:00", ""},

		{"2021-06-13", "2021-06-19 23:59:59"},
		{"2021-06-14", "2021-06-19 23:59:59"},
		{"2021-06-18", "2021-06-19 23:59:59"},
		{"2021-07-17", "2021-07-17 23:59:59"},
		{"2021-07-18", "2021-07-24 23:59:59"},
	}

	for index, test := range tests {
		c := Parse(test.input).EndOfWeek()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_StartOfDay(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2020-01-15 00:00:00"},
		{"2020-02-15 00:00:00", "2020-02-15 00:00:00"},
		{"2020-03-15 00:00:00", "2020-03-15 00:00:00"},
		{"2020-04-15 23:59:59", "2020-04-15 00:00:00"},
		{"2020-05-15 23:59:59", "2020-05-15 00:00:00"},
		{"2020-06-15 23:59:59", "2020-06-15 00:00:00"},
		{"2020-07-15 23:59:59", "2020-07-15 00:00:00"},
		{"2020-08-15 13:14:15", "2020-08-15 00:00:00"},
		{"2020-09-15 13:14:15", "2020-09-15 00:00:00"},
		{"2020-10-15", "2020-10-15 00:00:00"},
		{"2020-11-15", "2020-11-15 00:00:00"},
		{"2020-12-15", "2020-12-15 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).StartOfDay()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_EndOfDay(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2020-01-15 23:59:59"},
		{"2020-02-15 00:00:00", "2020-02-15 23:59:59"},
		{"2020-03-15 00:00:00", "2020-03-15 23:59:59"},
		{"2020-04-15 23:59:59", "2020-04-15 23:59:59"},
		{"2020-05-15 23:59:59", "2020-05-15 23:59:59"},
		{"2020-06-15 23:59:59", "2020-06-15 23:59:59"},
		{"2020-07-15 23:59:59", "2020-07-15 23:59:59"},
		{"2020-08-15 13:14:15", "2020-08-15 23:59:59"},
		{"2020-09-15 13:14:15", "2020-09-15 23:59:59"},
		{"2020-10-15", "2020-10-15 23:59:59"},
		{"2020-11-15", "2020-11-15 23:59:59"},
		{"2020-12-15", "2020-12-15 23:59:59"},
	}

	for index, test := range tests {
		c := Parse(test.input).EndOfDay()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_StartOfHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2020-01-15 00:00:00"},
		{"2020-02-15 00:00:00", "2020-02-15 00:00:00"},
		{"2020-03-15 00:00:00", "2020-03-15 00:00:00"},
		{"2020-04-15 23:59:59", "2020-04-15 23:00:00"},
		{"2020-05-15 23:59:59", "2020-05-15 23:00:00"},
		{"2020-06-15 23:59:59", "2020-06-15 23:00:00"},
		{"2020-07-15 23:59:59", "2020-07-15 23:00:00"},
		{"2020-08-15 13:14:15", "2020-08-15 13:00:00"},
		{"2020-09-15 13:14:15", "2020-09-15 13:00:00"},
		{"2020-10-15", "2020-10-15 00:00:00"},
		{"2020-11-15", "2020-11-15 00:00:00"},
		{"2020-12-15", "2020-12-15 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).StartOfHour()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_EndOfHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2020-01-15 00:59:59"},
		{"2020-02-15 00:00:00", "2020-02-15 00:59:59"},
		{"2020-03-15 00:00:00", "2020-03-15 00:59:59"},
		{"2020-04-15 23:59:59", "2020-04-15 23:59:59"},
		{"2020-05-15 23:59:59", "2020-05-15 23:59:59"},
		{"2020-06-15 23:59:59", "2020-06-15 23:59:59"},
		{"2020-07-15 23:59:59", "2020-07-15 23:59:59"},
		{"2020-08-15 13:14:15", "2020-08-15 13:59:59"},
		{"2020-09-15 13:14:15", "2020-09-15 13:59:59"},
		{"2020-10-15", "2020-10-15 00:59:59"},
		{"2020-11-15", "2020-11-15 00:59:59"},
		{"2020-12-15", "2020-12-15 00:59:59"},
	}

	for index, test := range tests {
		c := Parse(test.input).EndOfHour()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_StartOfMinute(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2020-01-15 00:00:00"},
		{"2020-02-15 00:00:00", "2020-02-15 00:00:00"},
		{"2020-03-15 00:00:00", "2020-03-15 00:00:00"},
		{"2020-04-15 23:59:59", "2020-04-15 23:59:00"},
		{"2020-05-15 23:59:59", "2020-05-15 23:59:00"},
		{"2020-06-15 23:59:59", "2020-06-15 23:59:00"},
		{"2020-07-15 23:59:59", "2020-07-15 23:59:00"},
		{"2020-08-15 13:14:15", "2020-08-15 13:14:00"},
		{"2020-09-15 13:14:15", "2020-09-15 13:14:00"},
		{"2020-10-15", "2020-10-15 00:00:00"},
		{"2020-11-15", "2020-11-15 00:00:00"},
		{"2020-12-15", "2020-12-15 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).StartOfMinute()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_EndOfMinute(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2020-01-15 00:00:59"},
		{"2020-02-15 00:00:00", "2020-02-15 00:00:59"},
		{"2020-03-15 00:00:00", "2020-03-15 00:00:59"},
		{"2020-04-15 23:59:59", "2020-04-15 23:59:59"},
		{"2020-05-15 23:59:59", "2020-05-15 23:59:59"},
		{"2020-06-15 23:59:59", "2020-06-15 23:59:59"},
		{"2020-07-15 23:59:59", "2020-07-15 23:59:59"},
		{"2020-08-15 13:14:15", "2020-08-15 13:14:59"},
		{"2020-09-15 13:14:15", "2020-09-15 13:14:59"},
		{"2020-10-15", "2020-10-15 00:00:59"},
		{"2020-11-15", "2020-11-15 00:00:59"},
		{"2020-12-15", "2020-12-15 00:00:59"},
	}

	for index, test := range tests {
		c := Parse(test.input).EndOfMinute()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_StartOfSecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00.123", "2020-01-15 00:00:00.0"},
		{"2020-02-15 00:00:00.123", "2020-02-15 00:00:00.0"},
		{"2020-03-15 00:00:00.123", "2020-03-15 00:00:00.0"},
		{"2020-04-15 23:59:59.123", "2020-04-15 23:59:59.0"},
		{"2020-05-15 23:59:59.123", "2020-05-15 23:59:59.0"},
		{"2020-06-15 23:59:59.123", "2020-06-15 23:59:59.0"},
		{"2020-07-15 23:59:59.123", "2020-07-15 23:59:59.0"},
		{"2020-08-15 13:14:15.123", "2020-08-15 13:14:15.0"},
		{"2020-09-15 13:14:15.123", "2020-09-15 13:14:15.0"},
		{"2020-10-15", "2020-10-15 00:00:00.0"},
		{"2020-11-15", "2020-11-15 00:00:00.0"},
		{"2020-12-15", "2020-12-15 00:00:00.0"},
	}

	for index, test := range tests {
		c := Parse(test.input).StartOfSecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Format("Y-m-d H:i:s.u"), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_EndOfSecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00.123", "2020-01-15 00:00:00.999"},
		{"2020-02-15 00:00:00.123", "2020-02-15 00:00:00.999"},
		{"2020-03-15 00:00:00.123", "2020-03-15 00:00:00.999"},
		{"2020-04-15 23:59:59.123", "2020-04-15 23:59:59.999"},
		{"2020-05-15 23:59:59.123", "2020-05-15 23:59:59.999"},
		{"2020-06-15 23:59:59.123", "2020-06-15 23:59:59.999"},
		{"2020-07-15 23:59:59.123", "2020-07-15 23:59:59.999"},
		{"2020-08-15 13:14:15.123", "2020-08-15 13:14:15.999"},
		{"2020-09-15 13:14:15.123", "2020-09-15 13:14:15.999"},
		{"2020-10-15", "2020-10-15 00:00:00.999"},
		{"2020-11-15", "2020-11-15 00:00:00.999"},
		{"2020-12-15", "2020-12-15 00:00:00.999"},
	}

	for index, test := range tests {
		c := Parse(test.input).EndOfSecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Format("Y-m-d H:i:s.u"), "Current test index is "+strconv.Itoa(index))
	}
}
