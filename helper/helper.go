package helper

import (
	"fmt"
	"goroutine-optimize/errs"
	"reflect"
)

type ResponseMessage struct {
	Error bool
	Code  int
	Data  any
}

func NewResponseMessage(err bool, code int, data any) *ResponseMessage {
	return &ResponseMessage{Error: err, Code: code, Data: data}
}

func IsValid(data interface{}) *errs.AppErr {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("required") == "true" && reflect.ValueOf(data).Field(i).Interface() == "" {
			return errs.NewValidationError(fmt.Sprintf("field %s cannot be empty", field.Name))
		}

		// lengthName := reflect.ValueOf(data).Field(i).Interface().(string) // .(string) di akhir digunakan untuk konversi tipe data pada tipe data interface kosong
		// fmt.Println()

		// lengthString, _ := strconv.Atoi(field.Tag.Get("max"))

		// // if disini untuk memvalidasi apakah panjang lebih dari tag max
		// if int(len(reflect.ValueOf(data).Field(i).Interface().(string))) > lengthString {
		// 	return errs.NewValidationError(fmt.Sprintf("field %s cannot more than 10", field.Name))
		// }
	}

	return nil
}