package xbyte

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func DecodeGzipBytes(meta []byte) ([]byte, error) {
	b := bytes.Buffer{}
	b.Write(meta)
	r, _ := gzip.NewReader(&b)
	defer r.Close()
	datas, readErr := ioutil.ReadAll(r)

	if readErr != nil {
		return nil, readErr
	}

	return datas, nil
}

func EncodeGzipBytes(meta []byte) []byte {
	b := bytes.Buffer{}
	w := gzip.NewWriter(&b)
	defer w.Close()

	_, _ = w.Write(meta)
	_ = w.Flush()

	return b.Bytes()
}
