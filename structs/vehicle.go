package structs

import (
	"math/rand"

	"github.com/ggmolly/go-json-benchmarks/utils"
)

const (
	// Try to reduce the number of enhancements if you don't have enough memory
	enhancementCount = 500
)

type Vehicle struct {
	Make          string     `json:"make"`
	Model         string     `json:"model"`
	Year          int        `json:"year"`
	StillProduced bool       `json:"still_produced"`
	Dimensions    [3]float64 `json:"dimensions"`
	Engine        *Engine    `json:"engine"`
}

type Engine struct {
	Manufacturer string             `json:"manufacturer"`
	Model        string             `json:"model"`
	Displacement float64            `json:"displacement"`
	Power        float64            `json:"power"`
	Torque       float64            `json:"torque"`
	FuelType     string             `json:"fuel_type"`
	Options      *[]string          `json:"options"`
	Enhancements *map[string]string `json:"enhancements"`
}

func RandomVehicle() Vehicle {
	return Vehicle{
		Make:          utils.RandomString(16),
		Model:         utils.RandomString(16),
		Year:          rand.Intn(2020) + 1900,
		StillProduced: rand.Intn(2) == 0,
		Dimensions:    [3]float64{rand.Float64() * 100, rand.Float64() * 100, rand.Float64() * 100},
		Engine:        RandomEngine(),
	}
}

func RandomEngine() *Engine {
	enhancements := make(map[string]string, enhancementCount)
	for i := 0; i < enhancementCount; i++ {
		enhancements[utils.RandomString(30)] = utils.RandomString(25)
	}
	return &Engine{
		Manufacturer: utils.RandomString(16),
		Model:        utils.RandomString(16),
		Displacement: rand.Float64() * 5.0,
		Power:        rand.Float64() * 1000,
		Torque:       rand.Float64() * 1000,
		FuelType:     utils.RandomString(16),
		Options:      &[]string{utils.RandomString(16), utils.RandomString(16), utils.RandomString(16)},
		Enhancements: &enhancements,
	}
}
