package util

import "strings"

var envArr = map[string]string{
	"production":  "production",
	"sandbox":     "sandbox",
	"development": "development",
	"prod":        "production",
	"sndbx":       "sandbox",
	"dev":         "development",
	"prd":         "production",
	"local":       "local",
}

func EnvironmentTransform(s string) string {
	v, ok := envArr[strings.ToLower(strings.Trim(s, " "))]

	if !ok {
		return ""
	}

	return v
}
