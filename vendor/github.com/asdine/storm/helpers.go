package storm

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

// toBytes turns an interface into a slice of bytes
func toBytes(key interface{}) ([]byte, error) {
	if key == nil {
		return nil, nil
	}
	if k, ok := key.([]byte); ok {
		return k, nil
	}
	if k, ok := key.(string); ok {
		return []byte(k), nil
	}
	if k, ok := key.(fmt.Stringer); ok {
		return []byte(k.String()), nil
	}
	if k, ok := key.(json.Marshaler); ok {
		return k.MarshalJSON()
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
