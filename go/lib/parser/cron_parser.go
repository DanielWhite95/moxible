package cron_parser

import (
	"time"
	"strings"
)

func checkLine(cronLine string) {
	cronFields := strings.Fields(cronLine)
	if len(cronFields) < 6 {
		return false
	}
	cronTime := cronFields[1:5] 
	cronCmd := cronFields[6:]
}

func checkMin(min string) {
}

func checkHrs(min string) {
}

func checkDay(min string) {
}

func checkMon(min string) {
}

func checkWeekDay(min string) {
}
