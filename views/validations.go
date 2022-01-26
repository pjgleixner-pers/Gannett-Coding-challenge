package views

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
)

var (
	// ErrInvalidItemId : Description about this
	ErrInvalidItemId = errors.New("invalid Item id")
	// ErrInvalidName : Description about this
	ErrInvalidName = errors.New("invalid name")
	// ErrInvalidUnitPrice : Description about this
	ErrInvalidUnitPrice = errors.New("invalid unit price")
)

type ItemValidation interface {
	Validate(r *http.Request) error
}

func (item Item) Validate(r *http.Request) error {
	if len(item.ID) != 19 {
		return ErrInvalidItemId
	}
	chars := []rune(item.ID)
	if !govalidator.IsAlphanumeric(string(chars[0:3])) {
		return ErrInvalidItemId
	}

	if !govalidator.IsAlphanumeric(string(chars[5:8])) {
		return ErrInvalidItemId
	}

	if !govalidator.IsAlphanumeric(string(chars[10:13])) {
		return ErrInvalidItemId
	}

	if !govalidator.IsAlphanumeric(string(chars[15:18])) {
		return ErrInvalidItemId
	}

	if govalidator.IsNull(item.Name) {
		return ErrInvalidName
	}

	if !govalidator.IsAlphanumeric(item.Name) {
		return ErrInvalidName
	}

	if !govalidator.IsFloat(item.Price) {
		return ErrInvalidUnitPrice
	}

	return nil
}

func Validate(r *http.Request, v ItemValidation) error {
	return v.Validate(r)
}