package main

import "time"

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

var monthToName = map[time.Month]string{
	1:  "Sausio",
	2:  "Vasario",
	3:  "Kovo",
	4:  "Balandžio",
	5:  "Gegužės",
	6:  "Birželio",
	7:  "Liepos",
	8:  "Rugpjūčio",
	9:  "Rugsėjo",
	10: "Spalio",
	11: "Lapkričio",
	12: "Gruodžio",
}

func TemperatureColor(temperature float32) string {
	switch {
	case temperature <= 12:
		return blue
	case temperature < 20:
		return cyan
	case temperature < 25:
		return green
	default: // 25+
		return red
	}
}

func GetRainDescription(totalParticipation float32) string {
	switch {
	case totalParticipation == 0:
		return "nelis"
	case totalParticipation <= 1:
		return "mažas lietus"
	case totalParticipation <= 2:
		return "vidutinis lietus"
	default: // 2+
		return "smarkus lietus"
	}
}

func MapMonthsToLithuanian(month time.Month) string {
	if monthName, ok := monthToName[month]; ok {
		return monthName
	}
	return "Nežinomas mėnuo"
}
