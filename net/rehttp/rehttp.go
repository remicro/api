package rehttp

import "github.com/remicro/api/serialization"

type Before func(b Builder, path string, body []byte)

type Response interface {
	Status() (code int)
	Decoded() (decoded interface{})
	ContentType() (contentType ContentType)
	Header(key string) (values []string)
	Body() (data []byte)
	Error() (err error)
}

type Builder interface {
	Address(address string) Builder
	GET(path string) Builder
	POST(path string) Builder
	PUT(path string) Builder
	DELETE(path string) Builder
	PATCH(path string) Builder
	OPTIONS(path string) Builder
	QueryParam(key, value string) Builder
	ToEncode(object interface{}) Builder
	ToDecode(object interface{}) Builder
	ContentType(contentType ContentType) Builder
	Header(key, value string) Builder
	Cookie(key string, value []byte) Builder
	Encoder(encoder serialization.Encoder) Builder
	Decoder(decoder serialization.Decoder) Builder
	Before(fn Before) Builder
	Go() (response Response, err error)
}

type Factory interface {
	To(address string) Builder
}

