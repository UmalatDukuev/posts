package utils

func GetStringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
