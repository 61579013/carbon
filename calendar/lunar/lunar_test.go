package lunar

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSolarToLunar(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(2024, 1, 21, 0, 0, 0, 0, time.Local))},
			want: "2023-12-11 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2023, 3, 2, 0, 0, 0, 0, time.Local))},
			want: "2023-02-11 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2023, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "2023-02-11 00:00:00",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().String(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunarToGregorian(t *testing.T) {
	type args struct {
		l Lunar
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewLunar(2023, 12, 11, 0, 0, 0, false)},
			want: "2024-01-21 00:00:00",
		},
		{
			args: args{NewLunar(2023, 2, 11, 0, 0, 0, false)},
			want: "2023-03-02 00:00:00",
		},
		{
			args: args{NewLunar(2023, 2, 11, 0, 0, 0, true)},
			want: "2023-04-01 00:00:00",
		},
		{
			args: args{NewLunar(1800, 2, 11, 0, 0, 0, false)},
			want: "",
		},
		{
			args: args{NewLunar(2500, 2, 11, 0, 0, 0, false)},
			want: "",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.l).ToGregorian().String(), "args{%v}", tt.args.l)
		})
	}
}

func TestLunar_Animal(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: "牛",
		},
		{
			args: args{NewGregorian(time.Date(2010, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "虎",
		},
		{
			args: args{NewGregorian(time.Date(2011, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "兔",
		},
		{
			args: args{NewGregorian(time.Date(2012, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "龙",
		},
		{
			args: args{NewGregorian(time.Date(2013, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "蛇",
		},
		{
			args: args{NewGregorian(time.Date(2014, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "马",
		},
		{
			args: args{NewGregorian(time.Date(2015, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "羊",
		},
		{
			args: args{NewGregorian(time.Date(2016, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "猴",
		},
		{
			args: args{NewGregorian(time.Date(2017, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "鸡",
		},
		{
			args: args{NewGregorian(time.Date(2018, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "狗",
		},
		{
			args: args{NewGregorian(time.Date(2019, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "猪",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 23, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			args: args{NewGregorian(time.Date(2020, 6, 21, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			args: args{NewGregorian(time.Date(2021, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "牛",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Animal(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Festival(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{NewGregorian(time.Date(2021, 2, 12, 0, 0, 0, 0, time.Local))},
			want: "春节",
		},
		{
			args: args{NewGregorian(time.Date(2021, 2, 26, 0, 0, 0, 0, time.Local))},
			want: "元宵节",
		},
		{
			args: args{NewGregorian(time.Date(2021, 6, 14, 0, 0, 0, 0, time.Local))},
			want: "端午节",
		},
		{
			args: args{NewGregorian(time.Date(2021, 8, 14, 0, 0, 0, 0, time.Local))},
			want: "七夕节",
		},
		{
			args: args{NewGregorian(time.Date(2021, 8, 22, 0, 0, 0, 0, time.Local))},
			want: "中元节",
		},
		{
			args: args{NewGregorian(time.Date(2021, 9, 21, 0, 0, 0, 0, time.Local))},
			want: "中秋节",
		},
		{
			args: args{NewGregorian(time.Date(2021, 10, 14, 0, 0, 0, 0, time.Local))},
			want: "重阳节",
		},
		{
			args: args{NewGregorian(time.Date(2021, 11, 5, 0, 0, 0, 0, time.Local))},
			want: "寒衣节",
		},
		{
			args: args{NewGregorian(time.Date(2021, 11, 19, 0, 0, 0, 0, time.Local))},
			want: "下元节",
		},
		{
			args: args{NewGregorian(time.Date(2022, 1, 10, 0, 0, 0, 0, time.Local))},
			want: "腊八节",
		},
		{
			args: args{NewGregorian(time.Date(2022, 1, 25, 0, 0, 0, 0, time.Local))},
			want: "小年",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Festival(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Year(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: 2020,
		},
		{
			args: args{NewGregorian(time.Date(2021, 5, 1, 0, 0, 0, 0, time.Local))},
			want: 2021,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Year(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Month(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2021, 3, 5, 0, 0, 0, 0, time.Local))},
			want: 1,
		},
		{
			args: args{NewGregorian(time.Date(2021, 4, 5, 0, 0, 0, 0, time.Local))},
			want: 2,
		},
		{
			args: args{NewGregorian(time.Date(2021, 5, 5, 0, 0, 0, 0, time.Local))},
			want: 3,
		},
		{
			args: args{NewGregorian(time.Date(2021, 6, 5, 0, 0, 0, 0, time.Local))},
			want: 4,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 5, 0, 0, 0, 0, time.Local))},
			want: 5,
		},
		{
			args: args{NewGregorian(time.Date(2021, 8, 5, 0, 0, 0, 0, time.Local))},
			want: 6,
		},
		{
			args: args{NewGregorian(time.Date(2021, 9, 5, 0, 0, 0, 0, time.Local))},
			want: 7,
		},
		{
			args: args{NewGregorian(time.Date(2021, 10, 5, 0, 0, 0, 0, time.Local))},
			want: 8,
		},
		{
			args: args{NewGregorian(time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local))},
			want: 9,
		},
		{
			args: args{NewGregorian(time.Date(2021, 11, 5, 0, 0, 0, 0, time.Local))},
			want: 10,
		},
		{
			args: args{NewGregorian(time.Date(2022, 12, 5, 0, 0, 0, 0, time.Local))},
			want: 11,
		},
		{
			args: args{NewGregorian(time.Date(2022, 1, 5, 0, 0, 0, 0, time.Local))},
			want: 12,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Month(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_LeapMonth(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: 4,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().LeapMonth(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Day(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 19, 0, 0, 0, 0, time.Local))},
			want: 1,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 20, 0, 0, 0, 0, time.Local))},
			want: 2,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 21, 0, 0, 0, 0, time.Local))},
			want: 3,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 22, 0, 0, 0, 0, time.Local))},
			want: 4,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Day(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToYearString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零",
		},
		{
			args: args{NewGregorian(time.Date(2021, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二一",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToYearString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToMonthString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{NewGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "腊月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "正月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "二月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "三月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: "闰四月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "闰四月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "闰四月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "五月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "六月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "七月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "八月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "九月",
		},
		{
			args: args{NewGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "十月",
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "十一月",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToMonthString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToDayString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{NewGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "初七",
		},
		{
			args: args{NewGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "初八",
		},
		{
			args: args{NewGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "初八",
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "初九",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "初九",
		},
		{
			args: args{NewGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "初十",
		},
		{
			args: args{NewGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "十一",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "十二",
		},
		{
			args: args{NewGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "十四",
		},
		{
			args: args{NewGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "十五",
		},
		{
			args: args{NewGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "十六",
		},
		{
			args: args{NewGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "十七",
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local))},
			want: "二十",
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 5, 0, 0, 0, 0, time.Local))},
			want: "廿二",
		},
		{
			args: args{NewGregorian(time.Date(2021, 4, 11, 0, 0, 0, 0, time.Local))},
			want: "三十",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToDayString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_String(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{NewGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "2019-12-07 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-01-08 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-02-08 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-03-09 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-04-09 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-04-10 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-05-11 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-06-12 00:00:00",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().String(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToDateString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{NewGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "二零一九年腊月初七",
		},
		{
			args: args{NewGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年正月初八",
		},
		{
			args: args{NewGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年二月初八",
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年三月初九",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年闰四月初九",
		},
		{
			args: args{NewGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年闰四月初十",
		},
		{
			args: args{NewGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年五月十一",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年六月十二",
		},
		{
			args: args{NewGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年七月十四",
		},
		{
			args: args{NewGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年八月十五",
		},
		{
			args: args{NewGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年九月十六",
		},
		{
			args: args{NewGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年十月十七",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToDateString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsLeapYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsLeapYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsLeapMonth(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsLeapMonth(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsRatYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsRatYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsOxYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsOxYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsTigerYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsTigerYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsRabbitYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2023, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsRabbitYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsDragonYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2024, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsDragonYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsSnakeYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2025, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsSnakeYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsHorseYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2026, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsHorseYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsGoatYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2027, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsGoatYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsMonkeyYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2028, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsMonkeyYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsRoosterYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2029, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsRoosterYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsDogYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2030, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsDogYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsPigYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2031, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsPigYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_DoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{NewGregorian(time.Date(2020, 1, 5, 23, 23, 45, 0, time.Local))},
			want: "子时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 59, 45, 0, time.Local))},
			want: "子时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 1, 0, 0, 0, time.Local))},
			want: "丑时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 3, 40, 0, 0, time.Local))},
			want: "寅时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 5, 0, 0, 0, time.Local))},
			want: "卯时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 7, 0, 0, 0, time.Local))},
			want: "辰时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 9, 0, 0, 0, time.Local))},
			want: "巳时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 11, 0, 0, 0, time.Local))},
			want: "午时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 13, 0, 0, 0, time.Local))},
			want: "未时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 15, 0, 0, 0, time.Local))},
			want: "申时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 17, 0, 0, 0, time.Local))},
			want: "酉时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 19, 0, 0, 0, time.Local))},
			want: "戌时",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 21, 0, 0, 0, time.Local))},
			want: "亥时",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().DoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsFirstDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 23, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 1, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsFirstDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsSecondDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 1, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsSecondDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsThirdDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 3, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsThirdDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsFourthDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 5, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsFourthDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsFifthDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 7, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsFifthDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsSixthDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 9, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsSixthDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsSeventhDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 11, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsSeventhDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsEighthDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 13, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsEighthDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsNinthDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 15, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsNinthDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsTenthDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 17, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsTenthDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsEleventhDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 19, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsEleventhDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsTwelfthDoubleHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 21, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsTwelfthDoubleHour(), "args{%v}", tt.args.g)
		})
	}
}
