package models

// Chicken is a bird
type Chicken interface {
	Cacarejar() string
}

// Duck is also a bird
type Duck interface {
	Quack() string
}

// Bird struct for all bird types
type Bird struct {
	Name string
}

// Cacarejar return the chicken sound
func (bird Bird) Cacarejar() string {
	return "Cócóricó..."
}

// Quack return the duck sound
func (bird Bird) Quack() string {
	return "Quack, quack...."
}
