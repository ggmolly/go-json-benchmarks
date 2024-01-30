package gojsonbenchmarks

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ggmolly/go-json-benchmarks/structs"
)

// TestStruct is a struct containing a type, and a function that returns an instance of that type
type TestStruct struct {
	Instance         reflect.Type
	PopulateFunction func() interface{}
}

var (
	Structs = map[string]TestStruct{
		"User": {
			Instance:         reflect.TypeOf(structs.User{}),
			PopulateFunction: func() interface{} { return structs.RandomUser() },
		},
		"Vehicle": {
			Instance:         reflect.TypeOf(structs.Vehicle{}),
			PopulateFunction: func() interface{} { return structs.RandomVehicle() },
		},
		"Family": {
			Instance:         reflect.TypeOf(structs.Family{}),
			PopulateFunction: func() interface{} { return structs.RandomFamily() },
		},
	}
)

func benchmarkType(b *testing.B, testStruct TestStruct, library Library) {
	// Create a slice of testStruct.PopulateFunction() with length b.N
	instances := make([]interface{}, b.N)
	for i := 0; i < b.N; i++ {
		instances[i] = testStruct.PopulateFunction()
	}

	// Marshal the entire list
	marshaled, err := library.Marshal(instances)
	if err != nil {
		b.Fatalf("Failed to marshal: %v", err)
	}

	// Unmarshal the entire list
	unmarshaled := make([]interface{}, b.N)
	err = library.Unmarshal(marshaled, &unmarshaled)
	if err != nil {
		b.Fatalf("Failed to unmarshal: %v", err)
	}
}

func BenchmarkAllStructs(b *testing.B) {
	for structName, testStruct := range Structs {
		for libraryName, library := range Libraries {
			b.Run(fmt.Sprintf("%s|%s", libraryName, structName), func(b *testing.B) {
				benchmarkType(b, testStruct, library)
			})
		}
	}
}
