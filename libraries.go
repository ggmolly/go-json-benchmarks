package gojsonbenchmarks

import (
	"encoding/json"

	sonic "github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
	segmentio "github.com/segmentio/encoding/json"
	jettison "github.com/wI2L/jettison"
)

type MarshalFunc func(interface{}) ([]byte, error)
type UnmarshalFunc func([]byte, interface{}) error

type Library struct {
	Marshal   func(interface{}) ([]byte, error)
	Unmarshal func([]byte, interface{}) error
}

var (
	Libraries = map[string]Library{
		"encoding/json": {
			Marshal:   json.Marshal,
			Unmarshal: json.Unmarshal,
		},
		"bytedance/sonic": {
			Marshal:   sonic.Marshal,
			Unmarshal: sonic.Unmarshal,
		},
		"goccy/gojson": {
			Marshal:   gojson.Marshal,
			Unmarshal: gojson.Unmarshal,
		},
		"segmentio/encoding": {
			Marshal:   segmentio.Marshal,
			Unmarshal: segmentio.Unmarshal,
		},
		"wI2L/jettison": {
			Marshal:   jettison.Marshal,
			Unmarshal: json.Unmarshal,
		},
	}
)
