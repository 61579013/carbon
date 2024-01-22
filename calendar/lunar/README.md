# Carbon

[![Carbon Release](https://img.shields.io/github/release/golang-module/carbon.svg)](https://github.com/golang-module/carbon/releases)
[![Go Test](https://github.com/golang-module/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/golang-module/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/carbon/v2)](https://goreportcard.com/report/github.com/golang-module/carbon/v2)
[![Go Coverage](https://codecov.io/gh/golang-module/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-module/carbon)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/golang-module/carbon/badges/download-count.svg)](https://goproxy.cn)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/carbon/v2)
[![License](https://img.shields.io/github/license/golang-module/carbon)](https://github.com/golang-module/carbon/blob/master/LICENSE)

English | [简体中文](README.cn.md) | [日本語](README.jp.md)

#### Introduction

A simple, semantic and developer-friendly golang package for time, has been included
by [awesome-go](https://github.com/avelino/awesome-go#date-and-time "awesome-go")

[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

[gitee.com/golang-module/carbon](https://gitee.com/golang-module/carbon "gitee.com/golang-module/carbon")

#### Chinese Lunar

> Currently only `200` years from `1900` to `2100` are supported

##### Convert Solar into Lunar

```go
// Get Lunar year of animal
carbon.Parse("2020-08-05 13:14:15").Lunar().Animal() // 鼠

// Get lunar festival
carbon.Parse("2021-02-12 13:14:15").Lunar().Festival() // 春节

// Get lunar year, month, day, hour, minute and second
carbon.Parse("2020-08-05 13:14:15").Lunar().DateTime() // 2020, 6, 16, 13, 14, 15
// Get lunar year, month and day
carbon.Parse("2020-08-05 13:14:15").Lunar().Date() // 2020, 6, 16
// Get lunar hour, minute and second
carbon.Parse("2020-08-05 13:14:15").Lunar().Time() // 13, 14, 15

// Get lunar year
carbon.Parse("2020-08-05 13:14:15").Lunar().Year() // 2020
// Get lunar month
carbon.Parse("2020-08-05 13:14:15").Lunar().Month() // 6
// Get lunar leap month
carbon.Parse("2020-08-05 13:14:15").Lunar().LeapMonth() // 4
// Get lunar day
carbon.Parse("2020-08-05 13:14:15").Lunar().Day() // 16
// Get lunar date as YYYY-MM-DD HH::ii::ss format string
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15").Lunar()) // 2020-06-16 13:14:15

// Get lunar year as string
carbon.Parse("2020-08-05 13:14:15").Lunar().ToYearString() // 二零二零
// Get lunar month as string
carbon.Parse("2020-08-05 13:14:15").Lunar().ToMonthString() // 六月
// Get lunar leap month as string
carbon.Parse("2020-04-23 13:14:15").Lunar().ToMonthString() // 闰四月
// Get lunar day as string
carbon.Parse("2020-08-05 13:14:15").Lunar().ToDayString() // 十六
// Get lunar date as string
carbon.Parse("2020-08-05 13:14:15").Lunar().ToDateString() // 二零二零年六月十六

// Whether is a lunar leap year
carbon.Parse("2020-08-05 13:14:15").Lunar().IsLeapYear() // true
// Whether is a lunar leap month
carbon.Parse("2020-08-05 13:14:15").Lunar().IsLeapMonth() // false

// Whether is a lunar year of the rat
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRatYear() // true
// Whether is a lunar year of the ox
carbon.Parse("2020-08-05 13:14:15").Lunar().IsOxYear() // false
// Whether is a lunar year of the tiger
carbon.Parse("2020-08-05 13:14:15").Lunar().IsTigerYear() // false
// Whether is a lunar year of the rabbit
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRabbitYear() // false
// Whether is a lunar year of the dragon
carbon.Parse("2020-08-05 13:14:15").Lunar().IsDragonYear() // false
// Whether is a lunar year of the snake
carbon.Parse("2020-08-05 13:14:15").Lunar().IsSnakeYear() // false
// Whether is a lunar year of the horse
carbon.Parse("2020-08-05 13:14:15").Lunar().IsHorseYear() // false
// Whether is a lunar year of the goat
carbon.Parse("2020-08-05 13:14:15").Lunar().IsGoatYear() // false
// Whether is a lunar year of the monkey
carbon.Parse("2020-08-05 13:14:15").Lunar().IsMonkeyYear() // false
// Whether is a lunar year of the rooster
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRoosterYear() // false
// Whether is a lunar year of the dog
carbon.Parse("2020-08-05 13:14:15").Lunar().IsDogYear() // false
// Whether is a lunar year of the dig
carbon.Parse("2020-08-05 13:14:15").Lunar().IsPigYear() // false

// Get lunar double-hour
carbon.Parse("2020-02-05 21:00:00").Lunar().DoubleHour() // 亥时

// Whether is FirstDoubleHour
carbon.Parse("2020-03-21 00:00:00").Lunar().IsFirstDoubleHour() // true
// Whether is SecondDoubleHour
carbon.Parse("2020-03-21 01:00:00").Lunar().IsSecondDoubleHour() // true
// Whether is ThirdDoubleHour
carbon.Parse("2020-03-21 03:00:00").Lunar().IsThirdDoubleHour() // true
// Whether is FourthDoubleHour
carbon.Parse("2020-03-21 05:00:00").Lunar().IsFourthDoubleHour() // true
// Whether is FifthDoubleHour
carbon.Parse("2020-03-21 07:00:00").Lunar().IsFifthDoubleHour() // true
// Whether is SixthDoubleHour
carbon.Parse("2020-03-21 09:00:00").Lunar().IsSixthDoubleHour() // true
// Whether is SeventhDoubleHour
carbon.Parse("2020-03-21 11:00:00").Lunar().IsSeventhDoubleHour() // true
// Whether is EighthDoubleHour
carbon.Parse("2020-03-21 13:00:00").Lunar().IsEighthDoubleHour() // true
// Whether is NinthDoubleHour
carbon.Parse("2020-03-21 15:00:00").Lunar().IsNinthDoubleHour() // true
// Whether is TenthDoubleHour
carbon.Parse("2020-03-21 17:00:00").Lunar().IsTenthDoubleHour() // true
// Whether is EleventhDoubleHour
carbon.Parse("2020-03-21 19:00:00").Lunar().IsEleventhDoubleHour() // true
// Whether is TwelfthDoubleHour
carbon.Parse("2020-03-21 21:00:00").Lunar().IsTwelfthDoubleHour() // true
```

##### Convert Lunar into Solar

```go
// Convert the Lunar Calendar December 11, 2023 to the Solar Calendar
carbon.CreateFromLunar(2023, 12, 11, 0, 0, 0, false).ToDateTimeString() // 2024-01-21 00:00:00
// Convert lunar calendar February 11, 2023 to Solar calendar
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, false).ToDateTimeString() // 2024-03-02 00:00:00
// Convert the Lunar Calendar Leap February 11, 2024 to the Solar Calendar
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, true).ToDateTimeString() // 2023-04-01 00:00:00
```