package controllers

import (
	"bytes"
	"io"
	"net/http"
)

type reusableReader struct {
	io.Reader
	readBuf *bytes.Buffer
	backBuf *bytes.Buffer
}

func ReusableReader(r io.Reader) io.Reader {
	readBuf := bytes.Buffer{}
	_, err := readBuf.ReadFrom(r)
	if err != nil {
		return nil
	}
	backBuf := bytes.Buffer{}

	return reusableReader{
		io.TeeReader(&readBuf, &backBuf),
		&readBuf,
		&backBuf,
	}
}

func (r reusableReader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	if err == io.EOF {
		r.reset()
	}
	return n, err
}

func (r reusableReader) reset() {
	_, err := io.Copy(r.readBuf, r.backBuf)
	if err != nil {
		return
	}
}

func ReuseBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = io.NopCloser(ReusableReader(r.Body))
		next.ServeHTTP(w, r)
	})
}
