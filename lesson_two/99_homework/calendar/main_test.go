package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

type Calendar struct {
	month   string
	quarter int
}

func NewCalendar(parsed time.Time) Calendar {

	calendar := Calendar{
		month:   strconv.Itoa(int(parsed.Month())),
		quarter: 0,
	}

	for i := 0; i < len(calendar.month); i++ {
		switch calendar.month {
		case "1", "2", "3":
			calendar.quarter = 1
		case "4", "5", "6":
			calendar.quarter = 2
		case "7", "8", "9":
			calendar.quarter = 3
		case "10", "11", "12":
			calendar.quarter = 4
		}
	}
	return calendar
}

func (c Calendar) CurrentQuarter() int {
	return c.quarter
}

func TestCurrentQuarter(t *testing.T) {
	cases := []struct {
		month   string
		quarter int
	}{
		{month: "01", quarter: 1},
		{month: "02", quarter: 1},
		{month: "03", quarter: 1},
		{month: "04", quarter: 2},
		{month: "05", quarter: 2},
		{month: "06", quarter: 2},
		{month: "07", quarter: 3},
		{month: "08", quarter: 3},
		{month: "09", quarter: 3},
		{month: "10", quarter: 4},
		{month: "11", quarter: 4},
		{month: "12", quarter: 4},
	}

	//TODO Реализовать Календарь

	for _, test := range cases {
		parsed, _ := time.Parse("2006-01-02", fmt.Sprintf("2015-%s-15", test.month))
		calendar := NewCalendar(parsed)
		actual := calendar.CurrentQuarter()
		if actual != test.quarter {
			t.Error("Month:", test.month,
				"Expected Quarter:", test.quarter,
				"Actual Quarter:", actual)
		}
	}
}
