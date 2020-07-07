package watcher

import (
	"strings"
	"time"
)

func CheckExcludedPaths(path string) bool {
	result := true
	items := []string{"", "/watcher/", "css", "/watcher", "/js/*", "/css/*", "/css/*filepath", "/js/*filepath", "/watcher/requests", "/js", "/watcher/responses", "/watcher/responses/:id", "/watcher/requests/:id"}
	for _, s := range items {
		if path == s {
			result = false
		}
	}
	return result
}

func UnixTimeToAmsterdam(rawTime int) string {
	loc, _ := time.LoadLocation("Europe/Amsterdam")
	timeInstance := time.Unix(int64(rawTime), 0)
	return timeInstance.In(loc).Format("15:04:05 Mon, 2 Jan 2006 ")
}

func formatJson(rawString string) string {
	if rawString == "" {
		return rawString
	}
	str := strings.ReplaceAll(rawString, "{", "{\n    ")
	str = strings.ReplaceAll(str, "}", "\n}")
	str = strings.ReplaceAll(str, "],", "],\n    ")
	return str
}