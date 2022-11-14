package helper

import (
	"reflect"
	"strings"
)

func GetFieldTag(object interface{}, fieldName string, tag string) string {
	//Get bson tag
	field, fieldErr := reflect.TypeOf(object).FieldByName(fieldName)

	if !fieldErr {
		//Do something
	}

	fullTag := field.Tag.Get(tag)
	splitTag := strings.Split(fullTag, ",")
	finalTag := splitTag[0]

	return finalTag
}
