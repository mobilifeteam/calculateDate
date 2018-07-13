package tngpdevgo

import (
	"strings"
)

func changeRemoveSlat(dateString string) (string, string, string) {
	day := strings.Split(dateString, "/")[0]
	month := strings.Split(dateString, "/")[1]
	year := strings.Split(dateString, "/")[2]
	return strings.TrimSpace(day), strings.TrimSpace(month), strings.TrimSpace(year)
}
