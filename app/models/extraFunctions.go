package models

func ReplaceNoneOrNanValueByEmptyString(val *string) {
	if *val == "None" || *val == "nan" {
		*val = ""
	}
}