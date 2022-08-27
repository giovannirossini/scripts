package models

import "errors"

// Realty is a struct that store data of a realty
type Realty struct {
	X     int    `json:"X"`
	Y     int    `json:"Y"`
	Name  string `json:"Name"`
	value int
}

// ErrorValueRealtyInvalid is an error that shows when the realty value is too high
var ErrorValueRealtyInvalid = errors.New("The realty value is not valid")

// ErrorValueRealty is an error that shows when the realty value is too short
var ErrorValueRealty = errors.New("The realty value is too high")

// SetValue function to set a min value for the realty
func (r *Realty) SetValue(value int) (err error) {
	err = nil
	if value < 50000 {
		err = ErrorValueRealtyInvalid
		return
	} else if value > 10000000 {
		err = ErrorValueRealty
		return
	}
	r.value = value
	return
}

// GetValue function to return the realty value
func (r *Realty) GetValue() int {
	return r.value
}
