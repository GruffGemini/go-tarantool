package datetime_test

import (
	"gopkg.in/vmihailenco/msgpack.v2"
)

type encoder = msgpack.Encoder
type decoder = msgpack.Decoder

func encodeUint(e *encoder, v uint64) error {
	return e.EncodeUint(uint(v))
}

func marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}