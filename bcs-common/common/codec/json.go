

package codec

import (
	"io"
	"reflect"

	"github.com/ugorji/go/codec"
)

var defaultJsonHandle = codec.JsonHandle{MapKeyAsString: true}

// DecJson json decode encapsulation
func DecJson(s []byte, v interface{}) error {
	dec := codec.NewDecoderBytes(s, &defaultJsonHandle)
	return dec.Decode(v)
}

// DecJsonReader json reader encapsulation
func DecJsonReader(s io.Reader, v interface{}) error {
	dec := codec.NewDecoder(s, &defaultJsonHandle)
	return dec.Decode(v)
}

// EncJson json encoder
func EncJson(v interface{}, s *[]byte) error {
	enc := codec.NewEncoderBytes(s, &defaultJsonHandle)
	return enc.Encode(v)
}

// EncJsonWriter json writer
func EncJsonWriter(v interface{}, s io.Writer) error {
	enc := codec.NewEncoder(s, &defaultJsonHandle)
	return enc.Encode(v)
}

func init() {
	defaultJsonHandle.MapType = reflect.TypeOf(map[string]interface{}(nil))
}
