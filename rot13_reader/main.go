package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(buf []byte) (n int, err error) {
	n, err = rot13.r.Read(buf)
	for i := 0; i < n; i++ {
		if 'A' <= buf[i] && buf[i] <= 'Z' {
			if buf[i]-'A' < 13 {
				buf[i] += 13
			} else {
				buf[i] -= 13
			}
		}
		if 'a' <= buf[i] && buf[i] <= 'z' {
			if buf[i]-'a' < 13 {
				buf[i] += 13
			} else {
				buf[i] -= 13
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
