package resources

import (
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
)

type SerialType string

const (
	JSON SerialType = "json"
	TOML SerialType = "toml"
)

// takes a series of bytes and attempts to construct an instance
// of the typed parameter using the byte content for values
func Deserialise[K any](data []byte, serialType SerialType) (*K, error) {
	var v K
	var err error
	
	switch serialType {
	case JSON:
		err = json.Unmarshal(data, &v)
	case TOML:
		_, err = toml.Decode(string(data), &v)
	default:
		return nil, fmt.Errorf("unknown serialisation type '%s'", serialType)
	}

	if err != nil {
		return nil, err
	}

	return &v, nil
}
