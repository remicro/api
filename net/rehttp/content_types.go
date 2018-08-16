package rehttp

type ContentType string

func (ct ContentType) String() string {
	return string(ct)
}

const (
	ContentTypeJson = ContentType("application/json")
	ConentTypeXML = ContentType("application/xml")
	ContentTypeHTML = ContentType("text/html")
)