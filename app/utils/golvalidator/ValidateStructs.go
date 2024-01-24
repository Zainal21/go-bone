package golvalidator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ValidateStructs(s interface{}, lang string) map[string][]string {
	var errors = make(map[string][]string)
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		config := v.Type().Field(i).Tag.Get("validate")
		fieldName := v.Type().Field(i).Tag.Get("json")
		fieldValue := v.Field(i).Interface()
		fieldError := validateField(config, fieldValue, fieldName, v, lang)
		if len(fieldError) > 0 {
			errors[fieldName] = fieldError
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return errors
}

func translateFieldName(fieldName string, lang string) string {
	var translations = map[string]string{
		"username":                "Username",
		"email":                   "Email",
		"Password":                "Kata Sandi",
		"oldPassword":             "Kata Sandi lama",
		"newPassword":             "Kata Sandi baru",
		"newPasswordConfirmation": "Konfirmasi kata sandi baru",
		"name":                    "Nama",
	}

	if translated, ok := translations[fieldName]; ok {
		return translated
	}

	return fieldName
}

func validateField(config string, fieldValue interface{}, fieldName string, refl reflect.Value, lang string) []string {
	var errors []string
	configArray := strings.Split(config, "|")
	v := reflect.ValueOf(fieldValue)
	dataType := ""
	for i := 0; i < len(configArray); i++ {
		validation := strings.Split(configArray[i], ":")
		switch validation[0] {
		case "nullable":
			if isEmptyValue(v) {
				return errors
			}
		default:
			newDataType, validatedField := validate(validation, v, fieldValue, fieldName, dataType, refl, lang)
			if newDataType != "" {
				dataType = newDataType
			}
			if validatedField != "" {
				errors = append(errors, validatedField)
			}
		}
	}

	return errors
}

func validate(validation []string, v reflect.Value, fieldValue interface{}, fieldName string, dataType string, refl reflect.Value, lang string) (string, string) {
	localType := dataType
	switch validation[0] {
	case "required":
		if isEmptyValue(v) {
			return localType, fmt.Sprintf(translate(lang, "%s wajib diisi.", "The %s field is required."), RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "alpha", "string":
		localType = "string"
		if !IsAlpha(v.String()) {
			return localType, fmt.Sprintf(translate(lang, "%s harus berupa string.", "The %s must only contain letters."), RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "numeric":
		localType = "string"
		if IsFloat(v.String()) {
			return localType, fmt.Sprintf(translate(lang, "%s harus berupa angka.", "The %s must be a number."), RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "alpha_num":
		if !IsAlphanumeric(v.String()) {
			return localType, fmt.Sprintf(translate(lang, "%s harus berupa huruf dan number.", "The %s must only contain letters and numbers"), RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "alpha_space":
		localType = "string"
		if !IsAlphaSpace(v.String()) {
			return localType, fmt.Sprintf("The %s must only contain letters, numbers, dashes and underscores.", RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "alpha_dash":
		localType = "string"
		if !IsAlphaDash(v.String()) {
			return localType, fmt.Sprintf("The %s must only contain letters, numbers, dashes and underscores.", RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "date":
		if !IsDate(v.String()) {
			return localType, fmt.Sprintf("The %s is not a valid date.", RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "email":
		if !IsEmail(v.String()) {
			return localType, fmt.Sprintf(translate(lang, "%s harus berupa alamat surel yang valid.", "The %s must be a valid email address"), RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "same":
		params := strings.Split(validation[1], ",")
		value2nd := refl.FieldByName(ToCamel(params[0])).Interface()
		if !IsSame(v.String(), value2nd.(string)) {
			return localType, fmt.Sprintf(
				translate(lang, "%s dan %s harus sama.", "The %s and %s must match."),
				RemoveUnderscore(translateFieldName(fieldName, lang)),
				RemoveUnderscore(translateFieldName(params[0], lang)),
			)
		}

	case "min":
		switch dataType {
		case "string":
			minCharacterLength, _ := strconv.Atoi(validation[1])
			textLength := len(v.String())

			if textLength < minCharacterLength {
				return localType, fmt.Sprintf("The %s must be at least %s characters.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}

		case "numeric":
			maxValue, _ := strconv.ParseFloat(validation[1], 64)
			floatValue, _ := ToFloat(fieldValue)

			if floatValue < maxValue {
				return localType, fmt.Sprintf("The %s must be at least %s.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}
		}

	case "max":
		switch dataType {
		case "string":
			maxCharacterLength, _ := strconv.Atoi(validation[1])
			textLength := len(v.String())

			if textLength > maxCharacterLength {
				return localType, fmt.Sprintf("The %s not be greater than %s characters.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}

		case "numeric":
			maxValue, _ := strconv.ParseFloat(validation[1], 64)
			floatValue, _ := ToFloat(fieldValue)

			if floatValue > maxValue {
				return localType, fmt.Sprintf("The %s not be greater than %s.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}
		}

	case "between":
		params := strings.Split(validation[1], ",")
		switch dataType {
		case "string":
			min, _ := strconv.Atoi(params[0])
			max, _ := strconv.Atoi(params[1])
			textLength := len(v.String())

			if !IsBetweenInt(textLength, min, max) {
				return localType, fmt.Sprintf("The %s must be between %s and %s characters.", RemoveUnderscore(translateFieldName(fieldName, lang)), params[0], params[1])
			}

		case "numeric":
			min, _ := strconv.ParseFloat(params[0], 64)
			max, _ := strconv.ParseFloat(params[1], 64)
			floatValue, _ := ToFloat(fieldValue)

			if !IsBetween(floatValue, min, max) {
				return localType, fmt.Sprintf("The %s must be between %s and %s.", RemoveUnderscore(translateFieldName(fieldName, lang)), params[0], params[1])
			}
		}

	case "digits":
		textLength := len(v.String())
		digits, _ := strconv.Atoi(validation[1])
		if textLength != digits {
			return localType, fmt.Sprintf(translate(lang, "%s harus %s karakter.", "The %s The %s must be %s digits."), RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
		}

	case "digits_between":
		params := strings.Split(validation[1], ",")
		min, _ := strconv.Atoi(params[0])
		max, _ := strconv.Atoi(params[1])
		textLength := len(v.String())

		if !IsBetweenInt(textLength, min, max) {
			return localType, fmt.Sprintf("The %s must be between %s and %s digits.", RemoveUnderscore(translateFieldName(fieldName, lang)), params[0], params[1])
		}

	case "lt":
		switch dataType {
		case "string":
			lessThanCharacterLength, _ := strconv.Atoi(validation[1])
			textLength := len(v.String())

			if !(textLength < lessThanCharacterLength) {
				return localType, fmt.Sprintf("The %s must be less than %s characters.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}

		case "numeric":
			lessThanValue, _ := strconv.ParseFloat(validation[1], 64)
			floatValue, _ := ToFloat(fieldValue)

			if !(floatValue < lessThanValue) {
				return localType, fmt.Sprintf("The %s must be less than %s.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}
		}

	case "gt":
		switch dataType {
		case "string":
			greaterThanCharacterLength, _ := strconv.Atoi(validation[1])
			textLength := len(v.String())

			if !(textLength > greaterThanCharacterLength) {
				return localType, fmt.Sprintf("The %s must be greater than %s characters.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}

		case "numeric":
			greaterThanValue, _ := strconv.ParseFloat(validation[1], 64)
			floatValue, _ := ToFloat(fieldValue)

			if !(floatValue > greaterThanValue) {
				return localType, fmt.Sprintf("The %s must be greater than %s.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}
		}

	case "lte":
		switch dataType {
		case "string":
			lessThanCharacterLength, _ := strconv.Atoi(validation[1])
			textLength := len(v.String())

			if !(textLength <= lessThanCharacterLength) {
				return localType, fmt.Sprintf("The %s must be less than %s characters.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}

		case "numeric":
			lessThanValue, _ := strconv.ParseFloat(validation[1], 64)
			floatValue, _ := ToFloat(fieldValue)

			if !(floatValue <= lessThanValue) {
				return localType, fmt.Sprintf("The %s must be less than %s.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}
		}

	case "gte":
		switch dataType {
		case "string":
			greaterThanCharacterLength, _ := strconv.Atoi(validation[1])
			textLength := len(v.String())

			if !(textLength >= greaterThanCharacterLength) {
				return localType, fmt.Sprintf("The %s must be greater than %s characters.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}

		case "numeric":
			greaterThanValue, _ := strconv.ParseFloat(validation[1], 64)
			floatValue, _ := ToFloat(fieldValue)

			if !(floatValue >= greaterThanValue) {
				return localType, fmt.Sprintf("The %s must be greater than %s.", RemoveUnderscore(translateFieldName(fieldName, lang)), validation[1])
			}
		}

	case "required_if":
		params := strings.Split(validation[1], ",")
		secondFieldValue := refl.FieldByName(ToCamel(params[0])).Interface()
		if v.String() != "" {
			return localType, ""
		}

		if secondFieldValue.(string) == params[1] {
			return localType, fmt.Sprintf("The %s field is required when %s is %s.", RemoveUnderscore(translateFieldName(fieldName, lang)), RemoveUnderscore(params[0]), params[1])
		}

	case "required_with":
		if v.String() != "" {
			return localType, ""
		}

		secondFieldValue := refl.FieldByName(ToCamel(validation[1])).Interface()
		if secondFieldValue != "" {
			return localType, fmt.Sprintf("The %s field is required when %s is present.", RemoveUnderscore(translateFieldName(fieldName, lang)), RemoveUnderscore(validation[1]))
		}

	case "ip":
		if !IsIP(v.String()) {
			return localType, fmt.Sprintf("The %s must be a valid IP address.", RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "ipv4":
		if !IsIPv4(v.String()) {
			return localType, fmt.Sprintf("The %s must be a valid IPv4 address.", RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "ipv6":
		if !IsIPv6(v.String()) {
			return localType, fmt.Sprintf("The %s must be a valid IPv6 address.", RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "url":
		if !IsURL(v.String()) {
			return localType, fmt.Sprintf("The %s format is invalid.", RemoveUnderscore(translateFieldName(fieldName, lang)))
		}

	case "credit_card":
		if !IsCreditCard(v.String()) {
			return localType, fmt.Sprintf("The %s must have a valid credit card number.", RemoveUnderscore(translateFieldName(fieldName, lang)))
		}
	}

	return localType, ""
}

func translate(lang, idMessage, enMessage string) string {
	switch lang {
	case "id":
		return idMessage
	default:
		return enMessage
	}
}
