package types

import "time"

// Flavour type to be stored in the Store
type Flavour struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Created_at time.Time `json:"created_at"`
}

// Request data type
type Create_flavour_request struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// New_falvour: create new flavour
func New_falvour(name string, price int) *Flavour {
	return &Flavour{
		Name:       name,
		Price:      price,
		Created_at: time.Now().UTC(),
	}
}
