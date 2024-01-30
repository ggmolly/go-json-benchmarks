package structs

import "github.com/ggmolly/go-json-benchmarks/utils"

const (
	// Try to reduce the number of enhancements if you don't have enough memory
	familyMember = 1000
)

type Family struct {
	LastName string `json:"last_name"`
	Members  []User `json:"members"`
}

func RandomFamily() Family {
	members := make([]User, familyMember)
	for i := 0; i < familyMember; i++ {
		members[i] = RandomUser()
	}
	return Family{
		LastName: utils.RandomString(10),
		Members:  members,
	}
}
