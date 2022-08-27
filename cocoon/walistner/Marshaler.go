//go:build ignore
// +build ignore

package main

import (
	"io"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type eventSourceMarshaler struct {
	runtime.JSONPb
}

type Marshaler interface {
	// Marshal marshals "v" into byte sequence.
	Marshal(v interface{}) ([]byte, error)
	// Unmarshal unmarshals "data" into "v".
	// "v" must be a pointer value.
	Unmarshal(data []byte, v interface{}) error
	// NewDecoder returns a Decoder which reads byte sequence from "r".
	NewDecoder(r io.Reader) Decoder
	// NewEncoder returns an Encoder which writes bytes sequence into "w".
	NewEncoder(w io.Writer) Encoder
	// ContentType returns the Content-Type which this marshaler is responsible for.
	// The parameter describes the type which is being marshalled, which can sometimes
	// affect the content type returned.
	ContentType(v interface{}) string
}
