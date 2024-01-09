package helpers

import (
	"encoding/json"
	"fmt"
	"html"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

func ReplaceInjectionString(input string) string {
	data := html.EscapeString(input)

	keywords := []string{
		"union",
		"database",
		"information_schema",
		"tabel_name",
		"columns",
		"'", "+",
		";DROP TABLE",
		"\\+",
		"1=1", "or", "OR",
	}
	for _, keyword := range keywords {
		data = strings.Replace(data, keyword, "", -1)
	}

	return data
}

func MapToJSON(m map[string]interface{}) string {
	jsonStr, _ := json.Marshal(m)
	return string(jsonStr)
}

func IsString(data interface{}) bool {
	_, isString := data.(string)
	return isString
}

func GetStringValue(data map[string]interface{}, key string) string {
	if value, ok := data[key]; ok {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return ""
}

func FormatDate(t time.Time, layout string) string {
	return t.Format(layout)
}

func ConvertTimeFormat(inputTime, inputFormat, outputFormat string) (string, error) {
	parsedTime, err := time.Parse(inputFormat, inputTime)
	if err != nil {
		return "", err
	}

	formattedTime := parsedTime.Format(outputFormat)
	return formattedTime, nil
}

func FormatDateToIndonesian(value string) (string, error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return "", err
	}

	parsedTime, err := time.Parse("02-01-2006", value)
	if err != nil {
		return "", err
	}

	parsedTime = parsedTime.In(loc)

	monthMap := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}

	englishMonth := parsedTime.Format("January")
	indonesianMonth, ok := monthMap[englishMonth]
	if !ok {
		return "", fmt.Errorf("unknown month: %s", englishMonth)
	}

	formattedDate := fmt.Sprintf("%d %s %d", parsedTime.Day(), indonesianMonth, parsedTime.Year())
	return formattedDate, nil
}

func GeneratePDFExportDukcapilFileName(NIK string) string {
	currentTime := time.Now()
	formattedDate := FormatDate(currentTime, "01022006")

	fileName := fmt.Sprintf("%s_%s_detail-nik.pdf", formattedDate, NIK)
	return fileName
}

func GetIntValue(data map[string]interface{}, key string) int {
	if value, ok := data[key]; ok {
		if num, ok := value.(float64); ok {
			return int(num)
		}
	}
	return 0
}
func GetTimeStrNow() string {
	currentTimestamps := time.Now()
	timeStr := currentTimestamps.Format(time.RFC3339)
	return timeStr
}

func Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func RemoveEmptyFields(payload interface{}) (interface{}, error) {
	value := reflect.ValueOf(payload)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("payload must be a struct or a pointer to a struct")
	}

	newValue := reflect.New(value.Type()).Elem()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		switch field.Kind() {
		case reflect.String:
			if field.String() != "" {
				newValue.Field(i).Set(field)
			}
		default:
			newValue.Field(i).Set(field)
		}
	}

	return newValue.Interface(), nil
}

func ConvertMapToJSONString(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	jsonString := string(jsonBytes)

	return jsonString, nil
}

func StructToMap(inputStruct interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	structValue := reflect.ValueOf(inputStruct)
	structType := structValue.Type()

	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structType.Field(i).Name
		fieldValue := structValue.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}

func StringReplaceToEmpty(value *string, defaultValue string) string {
	if value != nil {
		return *value
	}
	return defaultValue
}

func ValidateSortFields(sortField, sortType string, validSortFields map[string]string, defaultSortField, defaultSortType string) (string, string) {
	if _, valid := validSortFields[sortField]; !valid {
		sortField = defaultSortField
	}

	validSortTypes := map[string]bool{"asc": true, "desc": true}
	if _, valid := validSortTypes[sortType]; !valid {
		sortType = defaultSortType
	}

	return sortField, sortType
}

func GenerateRefId(length int) string {
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(randomString)
}
