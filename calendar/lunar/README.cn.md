# Carbon

[![Carbon Release](https://img.shields.io/github/release/golang-module/carbon.svg)](https://github.com/golang-module/carbon/releases)
[![Go Test](https://github.com/golang-module/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/golang-module/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/carbon/v2)](https://goreportcard.com/report/github.com/golang-module/carbon/v2)
[![Go Coverage](https://codecov.io/gh/golang-module/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-module/carbon)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/golang-module/carbon/badges/download-count.svg)](https://goproxy.cn)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/carbon/v2)
[![License](https://img.shields.io/github/license/golang-module/carbon)](https://github.com/golang-module/carbon/blob/master/LICENSE)

简体中文 | [English](README.md) | [日本語](README.jp.md)

一个轻量级、语义化、对开发者友好的 golang 时间处理库，支持链式调用，已被 [awesome-go-cn](https://github.com/yinggaozhen/awesome-go-cn#日期和时间 "awesome-go-cn") 收录

[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

[gitee.com/golang-module/carbon](https://gitee.com/golang-module/carbon "gitee.com/golang-module/carbon")

#### 中国农历

> 目前仅支持公元 `1900` 年至 `2100` 年的 `200` 年时间跨度

##### 将 `公历` 转换成 `农历`

```go
// 获取农历生肖
carbon.Parse("2020-08-05 13:14:15").Lunar().Animal() // 鼠

// 获取农历节日
carbon.Parse("2021-02-12 13:14:15").Lunar().Festival() // 春节

// 获取农历年月日时分秒
carbon.Parse("2020-08-05 13:14:15").Lunar().DateTime() // 2020, 6, 16, 13, 14, 15
// 获取农历年月日
carbon.Parse("2020-08-05 13:14:15").Lunar().Date() // 2020, 6, 16
// 获取农历时分秒
carbon.Parse("2020-08-05 13:14:15").Lunar().Time() // 13, 14, 15

// 获取农历年年份
carbon.Parse("2020-08-05 13:14:15").Lunar().Year() // 2020
// 获取农历月月份
carbon.Parse("2020-08-05 13:14:15").Lunar().Month() // 6
// 获取农历闰月月份
carbon.Parse("2020-08-05 13:14:15").Lunar().LeapMonth() // 4
// 获取农历日日期
carbon.Parse("2020-08-05 13:14:15").Lunar().Day() // 16
// 获取农历 YYYY-MM-DD HH::ii::ss 格式字符串
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15").Lunar()) // 2020-06-16 13:14:15

// 获取农历年字符串
carbon.Parse("2020-08-05 13:14:15").Lunar().ToYearString() // 二零二零
// 获取农历月字符串
carbon.Parse("2020-08-05 13:14:15").Lunar().ToMonthString() // 六月
// 获取农历闰月字符串
carbon.Parse("2020-04-23 13:14:15").Lunar().ToMonthString() // 闰四月
// 获取农历天字符串
carbon.Parse("2020-08-05 13:14:15").Lunar().ToDayString() // 十六
// 获取农历日期字符串
carbon.Parse("2020-08-05 13:14:15").Lunar().ToDateString() // 二零二零年六月十六

// 是否是农历闰年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsLeapYear() // true
// 是否是农历闰月
carbon.Parse("2020-08-05 13:14:15").Lunar().IsLeapMonth() // false

// 是否是鼠年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRatYear() // true
// 是否是牛年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsOxYear() // false
// 是否是虎年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsTigerYear() // false
// 是否是兔年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRabbitYear() // false
// 是否是龙年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsDragonYear() // false
// 是否是蛇年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsSnakeYear() // false
// 是否是马年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsHorseYear() // false
// 是否是羊年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsGoatYear() // false
// 是否是猴年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsMonkeyYear() // false
// 是否是鸡年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRoosterYear() // false
// 是否是狗年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsDogYear() // false
// 是否是猪年
carbon.Parse("2020-08-05 13:14:15").Lunar().IsPigYear() // false

// 获取农历时辰
carbon.Parse("2020-02-05 21:00:00").Lunar().DoubleHour() // 亥时

// 是否是子时
carbon.Parse("2020-03-21 00:00:00").Lunar().IsFirstDoubleHour() // true
// 是否是丑时
carbon.Parse("2020-03-21 01:00:00").Lunar().IsSecondDoubleHour() // true
// 是否是寅时
carbon.Parse("2020-03-21 03:00:00").Lunar().IsThirdDoubleHour() // true
// 是否是卯时
carbon.Parse("2020-03-21 05:00:00").Lunar().IsFourthDoubleHour() // true
// 是否是辰时
carbon.Parse("2020-03-21 07:00:00").Lunar().IsFifthDoubleHour() // true
// 是否是巳时
carbon.Parse("2020-03-21 09:00:00").Lunar().IsSixthDoubleHour() // true
// 是否是午时
carbon.Parse("2020-03-21 11:00:00").Lunar().IsSeventhDoubleHour() // true
// 是否是未时
carbon.Parse("2020-03-21 13:00:00").Lunar().IsEighthDoubleHour() // true
// 是否是申时
carbon.Parse("2020-03-21 15:00:00").Lunar().IsNinthDoubleHour() // true
// 是否是酉时
carbon.Parse("2020-03-21 17:00:00").Lunar().IsTenthDoubleHour() // true
// 是否是戌时
carbon.Parse("2020-03-21 19:00:00").Lunar().IsEleventhDoubleHour() // true
// 是否是亥时
carbon.Parse("2020-03-21 21:00:00").Lunar().IsTwelfthDoubleHour() // true
```

##### 将 `农历` 转化成 `公历`

```go
// 将农历 二零二三年腊月十一 转化为 公历
carbon.CreateFromLunar(2023, 12, 11, 0, 0, 0, false).ToDateTimeString() // 2024-01-21 00:00:00
// 将农历 二零二三年二月十一 转化为 公历
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, false).ToDateTimeString() // 2023-03-02 00:00:00
// 将农历 二零二三年闰二月十一 转化为 公历
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, true).ToDateTimeString() // 2023-04-01 00:00:00
```