package serialization

import "io"

type Encoder interface {
	Encode(object interface{}) (data []byte, err error)
}

type Decoder interface {
	Decode(object interface{}, data []byte) (err error)
}

type StreamDecoder interface {
	Decode(object interface{}) (err error)
}

type StreamEncoder interface {
	Encode(object interface{}) (err error)
}

type StreamFactory interface {
	Decode(at io.Reader) (StreamDecoder)
	Encode(at io.Writer) (StreamEncoder)
}
