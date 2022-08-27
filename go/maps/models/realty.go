package models

// Realty is a struct that store data of a realty
type Realty struct {
	X     int    `json:"X"`
	Y     int    `json:"Y"`
	Name  string `json:"Name"`
	value int
}

// SetValue function to set a min value for the realty
func (r *Realty) SetValue(value int) {
	r.value = value
}

// GetValue function to return the realty value
func (r *Realty) GetValue() int {
	return r.value
}
