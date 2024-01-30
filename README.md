# Go JSON Benchmarks

This project is a benchmark comparison of various Go libraries for JSON serialization and deserialization.

The goal is to provide a comprehensive comparison of their performance in different scenarios.

Only drop-in replacements for the standard library's `encoding/json` are considered.

Data integrity is verified, however, the time consumed by the verification is not included in the benchmarks.

## Results

Benchmarks were executed on a Ryzen 9 5900X with 48GB of RAM.

```
goos: linux
goarch: amd64
pkg: github.com/ggmolly/go-json-benchmarks
cpu: AMD Ryzen 9 5900X 12-Core Processor            
BenchmarkAllStructs/goccy/gojson|Vehicle-24                11607            510144 ns/op          438623 B/op       3609 allocs/op
BenchmarkAllStructs/segmentio/encoding|Vehicle-24          10381            557413 ns/op          463092 B/op       4613 allocs/op
BenchmarkAllStructs/wI2L/jettison|Vehicle-24                8779            662688 ns/op          432868 B/op       3606 allocs/op
BenchmarkAllStructs/encoding/json|Vehicle-24                8384            689001 ns/op          374542 B/op       4610 allocs/op
BenchmarkAllStructs/bytedance/sonic|Vehicle-24             13551            409650 ns/op          332513 B/op       2573 allocs/op
BenchmarkAllStructs/encoding/json|Family-24                 1864           3119586 ns/op         1777781 B/op      29023 allocs/op
BenchmarkAllStructs/bytedance/sonic|Family-24               3290           1727336 ns/op         1865884 B/op      16013 allocs/op
BenchmarkAllStructs/goccy/gojson|Family-24                  2272           2421286 ns/op         2705745 B/op      31014 allocs/op
BenchmarkAllStructs/segmentio/encoding|Family-24            1269           4120487 ns/op         7022062 B/op      35024 allocs/op
BenchmarkAllStructs/wI2L/jettison|Family-24                 1634           3209478 ns/op         2351772 B/op      29023 allocs/op
BenchmarkAllStructs/encoding/json|User-24                1799301              3269 ns/op            1785 B/op         30 allocs/op
BenchmarkAllStructs/bytedance/sonic|User-24              3058538              1763 ns/op            1686 B/op         17 allocs/op
BenchmarkAllStructs/goccy/gojson|User-24                 2184428              2502 ns/op            2556 B/op         31 allocs/op
BenchmarkAllStructs/segmentio/encoding|User-24           1617754              3875 ns/op            6990 B/op         36 allocs/op
BenchmarkAllStructs/wI2L/jettison|User-24                1741857              3303 ns/op            2267 B/op         30 allocs/op
PASS
ok      github.com/ggmolly/go-json-benchmarks   140.167s
```

## Running the benchmarks

```bash
git clone https://github.com/ggmolly/go-json-benchmarks.git
cd go-json-benchmarks
go test -bench . -benchmem -benchtime=5s
```

## Tested libraries

- [encoding/json](https://golang.org/pkg/encoding/json/)
- [bytedance/sonic](https://github.com/bytedance/sonic)
- [goccy/go-json](https://github.com/goccy/go-json)
- [segmentio/encoding/json](https://github.com/segmentio/encoding)
- [wI2L/jettison](https://github.com/wI2L/jettison)

## Contributing

I tried to make it as easy as possible to add new libraries and benchmarks.

### Adding a library

1. Add the library to `go.mod`
2. Add both `Marshal` and `Unmarshal` functions to `Libraries` in [libraries.go](libraries.go)

### Adding a struct

1. Create the file `structs/<struct_name>.go`
2. Add your struct definition
3. Add a function to randomly generate an instance of your struct to `Structs` in [structs.go](structs.go)
4. Add the struct to `Structs` in [structs_test.go](structs_test.go)

## Informations about the benchmarks

### Tested Structures

These structures were used for the benchmarks, they were chosen to represent common use cases in web applications.

#### User

```go
type User struct {
    ID       int64   `json:"id"`
    Username string  `json:"username"`
    Password string  `json:"password"`
    Email    string  `json:"email"`
    Phone    string  `json:"phone"`
    Active   bool    `json:"active"`
    Rating   float64 `json:"rating"`
}
```

#### Vehicle

```go
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
```

#### Family

```go
type Family struct {
    LastName string `json:"last_name"`
    Members  []User `json:"members"`
}
```

### Benchmark coverage

The following types are covered by the benchmarks:

- `string`
- `int`
- `int64`
- `float64`
- `bool`
- `[]string`
- `map[string]string`
- nested `structs`
- array of nested `structs`
- pointers
- maps (tested with `map[string]string`)