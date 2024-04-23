package validator

import (
	"errors"
	"reflect"
	"strings"
)

//теги полей структуры
const jsonTag = "json"
const requiredTag = "binding"

func ValidateHttpReq(requestStruct interface{}, tag string) (error) {
	var errStr []string

	t := reflect.TypeOf(requestStruct) //для работы с типами структуры
	v := reflect.ValueOf(requestStruct) //для работы со значениями структуры

	for i := 0; i < t.NumField(); i++ { //идём по полям структуры
		typeField := t.Field(i) 
		valField := v.Field(i)
		jsTag := typeField.Tag.Get(jsonTag) //значение тега json
		reqTag := typeField.Tag.Get(requiredTag) //значение тега validate

		f := valField.Interface()
		val := reflect.ValueOf(f) //значение переменной
		
		if typeField.Type.Kind().String() == "struct" {
			tt := val.Interface()
			
			if len(tag) != 0 && i == 0 {
				tag = tag+"." 
			}
			jsTag = tag+jsTag //Добавляем префикс структуры к полю

			err := ValidateHttpReq(tt, jsTag) //рекурсия для вложенных структур
			if err != nil {
				errStr = append(errStr, err.Error())
			}
		}
		if val.IsZero() && reqTag == "required" {
			if len(tag) != 0 && i == 0 {
				tag = tag+"."
			}
			errStr = append(errStr, tag+jsTag)
		} 
	}
	if len(errStr) != 0 {
		return errors.New(strings.Join(errStr, ", "))
	}
	return nil
}