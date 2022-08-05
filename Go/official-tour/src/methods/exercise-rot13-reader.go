package methods

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(buf []byte) (int, error) {
	length, err := rot13.r.Read(buf)
	if err != nil {
		return length, err
	}

	for i := 0; i < length; i++ {
		v := buf[i]
		switch {
		case 'a' <= v && v <= 'm':
			fallthrough
		case 'A' <= v && v <= 'M':
			buf[i] = v + 13
		case 'n' <= v && v <= 'z':
			fallthrough
		case 'N' <= v && v <= 'Z':
			buf[i] = v - 13
		}
	}
	return length, nil
}

func ExerciseRot13Reader() {
	s := strings.NewReader("Lbh pengxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println("")
}