package helper

import (
	"context"
	"encoding/base64"
	"fmt"
	"goroutine-optimize/errs"
	"math/rand"
	"reflect"
	"strconv"
	"sync"

	"golang.org/x/crypto/bcrypt"
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

		if field.Tag.Get("required") != "" && field.Tag.Get("required") == "true" && reflect.ValueOf(data).Field(i).Interface() == "" {
			return errs.NewValidationError(fmt.Sprintf("field %s cannot be empty", field.Name))
		}

		if field.Tag.Get("min") != "" {
			minLength, _ := strconv.Atoi(field.Tag.Get("min"))
			if int(len(reflect.ValueOf(data).Field(i).Interface().(string))) < minLength {
				return errs.NewValidationError(fmt.Sprintf("field %s cannot less than %s", field.Name, strconv.Itoa(minLength)))
			}
		}
	}
	return nil
}
func IsValidV2(data interface{}, err chan *errs.AppErr, ctx context.Context) {
	// * create wait group for waiting goroutine
	var wg sync.WaitGroup
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		wg.Add(1)

		// TODO : Validasi request client
		// ? validasi request client running on goroutine, so we must be waiting until goroutine done
		go func(i int) {
			defer wg.Done()

			// * why we use select ?. because we listen signal cancel from context.
			select {
			// TODO : if context send signal cancel. we cancel all goroutine.
			case <-ctx.Done():
				return
			default:
				// * this block for validasi all field from request client.
				if field.Tag.Get("required") != "" && field.Tag.Get("required") == "true" && reflect.ValueOf(data).Field(i).Interface() == "" {
					err <- errs.NewValidationError(fmt.Sprintf("field %s cannot be empty", field.Name))
					return
				}

				if field.Tag.Get("min") != "" {
					minLength, _ := strconv.Atoi(field.Tag.Get("min"))
					if int(len(reflect.ValueOf(data).Field(i).Interface().(string))) < minLength {
						err <- errs.NewValidationError(fmt.Sprintf("field %s cannot less than %s", field.Name, strconv.Itoa(minLength)))
						return
					}
				}
				return
			}
		}(i)
	}
	wg.Wait()

	// * if all field from request client done to validate, we send message End Of Line
	// * that for tell to goroutine which listent to channel to close channel. because theres no data to send in channel.
	err <- errs.NewUnexpectedError("End Of Line")

}

func BcryptPassword(passwordSalt string) string {
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordSalt), 8)
	return string(newPassword)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomStringV2() string {
	return base64.StdEncoding.EncodeToString([]byte("this is token"))
}
