package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(bytes []byte) (n int, err error) {
	byteA := []byte("A")[0]
	for i := range bytes {
		bytes[i] = byteA
	}
	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}
