package rehttp

type ContentType string

func (ct ContentType) String() string {
	return string(ct)
}

const (
	ContentTypeJson = "application/json"
	ConentTypeXML = "application/xml"
	ContentTypeHTML = "text/html"
)