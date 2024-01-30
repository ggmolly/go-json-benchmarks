package structs

import (
	"math/rand"

	"github.com/ggmolly/go-json-benchmarks/utils"
)

// A simple test struct
type User struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Active   bool    `json:"active"`
	Rating   float64 `json:"rating"`
}

// Fills a User struct with random data
func RandomUser() User {
	return User{
		ID:       rand.Int63(),
		Username: utils.RandomString(28),
		Password: utils.RandomString(28),
		Email:    utils.RandomString(28),
		Phone:    utils.RandomString(10),
		Active:   rand.Intn(2) == 0,
		Rating:   rand.Float64() * 5.0,
	}
}
