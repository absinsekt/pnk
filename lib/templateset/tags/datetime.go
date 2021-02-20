package tags

import (
	"fmt"
	"time"
)

var longMonthNames = []string{
	"",
	"январь",
	"февраль",
	"март",
	"апрель",
	"май",
	"июнь",
	"июль",
	"август",
	"сентябрь",
	"октябрь",
	"ноябрь",
	"декабрь",
}

var longMonthNamesAtDate = []string{
	"",
	"января",
	"февраля",
	"мара",
	"апреля",
	"мая",
	"июня",
	"июля",
	"августа",
	"сентября",
	"октября",
	"ноября",
	"декабря",
}

var datetimeFuncs = map[string]interface{}{
	"now": func(format string) string {
		return time.Now().Format(format)
	},
	"nowMiddle": func() string {
		now := time.Now()
		return fmt.Sprintf("%-0.2d %s, %d", now.Day(), longMonthNamesAtDate[now.Month()], now.Year())
	},
	"toShort": func(date time.Time) string {
		return date.Format("02.01.2006")
	},
}
