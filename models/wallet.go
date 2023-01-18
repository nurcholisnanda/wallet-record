package models

import (
	"time"

	validate "github.com/go-playground/validator/v10"
)

// datetimevalidation is a validation function that checks if the field passed in is of type time.Time and if it is not zero.
var datetimevalidation = func(fl validate.FieldLevel) bool {
	datetime, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}
	if datetime.IsZero() {
		return false
	}
	return true
}

// amountvalidation is a validation function that checks if the field passed in is of type float64 and if it is greater than 0.
var amountvalidation = func(fl validate.FieldLevel) bool {
	_, ok := fl.Field().Interface().(float64)
	if !ok {
		return false
	}
	return true
}

// The Validate function for the Record struct is defined to register the custom validation functions and then validate the struct.
func (r *Record) Validate() error {
	new := validate.New()
	err := new.RegisterValidation("datetimevalidation", datetimevalidation)
	if err != nil {
		return err
	}
	err = new.RegisterValidation("amountvalidation", amountvalidation)
	if err != nil {
		return err
	}
	err = new.Struct(r)
	if err != nil {
		return err
	}
	return err
}

// Record struct is defined with ID, Datetime and Amount fields.
type Record struct {
	ID       int       `gorm:"primary_key" json:"-"`
	Datetime time.Time `gorm:"type:datetime" json:"datetime" validate:"required,datetimevalidation"`
	Amount   float64   `gorm:"default:0" json:"amount" validate:"required,amountvalidation"`
}

// BasicResponse struct is defined with Success, Status and Message fields.
type BasicResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// The Validate function for the History struct is defined to register the custom validation function and then validate the struct.
func (h *History) Validate() error {
	new := validate.New()
	err := new.RegisterValidation("datetimevalidation", datetimevalidation)
	if err != nil {
		return err
	}
	err = new.Struct(h)
	if err != nil {
		return err
	}
	return err
}

// History struct is defined with StartDatetime and EndDatetime fields.
type History struct {
	StartDatetime time.Time `json:"startDatetime" validate:"required,datetimevalidation"`
	EndDatetime   time.Time `json:"endDatetime" validate:"required,datetimevalidation"`
}
