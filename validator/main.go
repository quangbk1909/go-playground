package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	v := validator.New()
	p := Person{
		Name:   "asfasdgasdgasdf",
		Age:    10,
		Gender: 1,
	}
	err := v.Var(p.Name, "required,max=5")
	if err != nil {
		ConvertToVinValidationErrors(err.(validator.ValidationErrors), p)
	}
}

type Person struct {
	Name   string `json:"name" validate:"required,max=5"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}

func ConvertToVinValidationErrors(errs validator.ValidationErrors, objectNeedValidate interface{}) {
	for _, e := range errs {
		var targetField reflect.StructField
		value := reflect.ValueOf(objectNeedValidate)
		field := reflect.TypeOf(objectNeedValidate)
		listNs := strings.Split(e.StructNamespace(), ".")
		for _, ns := range listNs[1:] {
			if value.Kind() == reflect.Ptr {
				value = value.Elem()
			}
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}

			regexNS := regexp.MustCompile(`^(?P<key>\w*)\[(?P<index>\d*)\]$`).FindStringSubmatch(ns)
			if regexNS != nil && len(regexNS) >= 3 {
				k := regexNS[1]
				i, _ := strconv.ParseInt(regexNS[2], 10, 64)

				if value.Kind() == reflect.Struct {
					value = value.FieldByName(k).Index(int(i))
					field = value.Type()
				}
			} else {
				targetField, _ = field.FieldByName(ns)
				if value.Kind() == reflect.Struct {
					value = value.FieldByName(ns)
					field = value.Type()
				}
			}
		}
		fmt.Println(targetField)

		//if targetField.Name != "" {
		//	vinErrors.Errors = append(vinErrors.Errors, FieldError{
		//		Field: targetField.Tag.Get("vin_validation_field"),
		//		Errors: []DetailError{
		//			{
		//				Code:    e.Tag(),
		//				Message: targetField.Tag.Get("vin_error_message_" + e.Tag()),
		//			},
		//		},
		//	})
		//}
	}
	return
}
