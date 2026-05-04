package xhttp

import (
	"net/http"
	"strings"
)

func IsGzipEncode(header http.Header) bool {
	value := header["Content-Encoding"]

	return value != nil && strings.EqualFold(value[0], "gzip")
}
