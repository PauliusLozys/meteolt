package main

import "time"

const (
	Gray   = "\033[37m"
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[97m"
)

func TemperatureColor(temperature float32) string {
	switch {
	case temperature <= 12:
		return Blue
	case temperature > 12 && temperature < 20:
		return Cyan
	case temperature >= 20 && temperature < 25:
		return Green
	case temperature >= 25:
		return Red
	default:
		return Reset
	}
}

func GetRainDescription(totalParticipation float32) string {
	switch {
	case totalParticipation == 0:
		return "nelis"
	case totalParticipation <= 1:
		return "mažas lietus"
	case totalParticipation > 1 && totalParticipation <= 2:
		return "vidutinis lietus"
	case totalParticipation > 2:
		return "smarkus lietus"
	}
	return ""
}

func MapMonthsToLithuanian(month time.Month) string {
	m := month - 1
	if m < 0 || m > 11 {
		return "Nežinomas mėnuo"
	}
	tt := []string{"Sausio", "Vasario", "Kovo", "Balandžio", "Gegužės", "Birželio", "Liepos", "Rugpjūčio", "Rugsėjo", "Spalio", "Lapkričio", "Gruodžio"}
	return tt[m]
}
