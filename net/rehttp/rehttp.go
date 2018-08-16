package rehttp

import "github.com/remicro/api/serialization"

type Response interface {
	Status() (code int)
	Decoded() (decoded interface{})
	ContentType() (contentType ContentType)
	Header(key string) (values []string)
	Body() (data []byte)
	Error() (err error)
}

type Builder interface {
	GET(url string) Builder
	POST(url string) Builder
	PUT(url string) Builder
	DELETE(url string) Builder
	PATCH(url string) Builder
	OPTIONS(url string) Builder
	QueryParam(key, value string) Builder
	ToEncode(object interface{}) Builder
	ToDecode(object interface{}) Builder
	ContentType(contentType ContentType) Builder
	Header(key, value string) Builder
	Cookie(key string, value []byte) Builder
	Encoder(encoder serialization.Encoder) Builder
	Decoder(decoder serialization.Decoder) Builder
	Go() (response Response, err error)
}

type Factory interface {
	To(address string) Builder
}

