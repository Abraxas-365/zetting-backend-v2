package models

import "strings"

type Format string

func (f Format) GetFormat(str string) Format {
	arr := strings.Split(str, ".")
	return Format(arr[len(arr)-1])
}
