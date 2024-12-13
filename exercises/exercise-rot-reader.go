package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := range b {
		if b[i] < 65 || b[i] > 122 {
			continue
		}
		switch {
			case b[i] >= 65 && b[i] <= 77:
				b[i] += 13
			case b[i] > 77 && b[i] <= 90:
				b[i] -= 13
			case b[i] >= 97 && b[i] <= 109:
				b[i] += 13
			case b[i] > 109 && b[i] <= 122:
				b[i] -= 13
		}
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

