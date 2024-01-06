package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!") // creates reader of passed arg

	b := make([]byte, 8)
	for {
		n, err := r.Read(b) // reads 8 chars from reader
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF { // if err == EOF end iterations
			break
		}
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r13 := Rot13Reader{s}
	io.Copy(os.Stdout, &r13)
}

type Rot13Reader struct {
	r io.Reader
}

func (rot13Reader Rot13Reader) Read(arr []byte) (n int, err error) {
	int, _ := rot13Reader.r.Read(arr)

	for i := 0; i < int; i++ {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			if (arr[i] + 13) > 'Z' {
				targetUnicodeNumber := 'A' + (arr[i] + 12 - 'Z')
				arr[i] = targetUnicodeNumber
			} else {
				arr[i] = arr[i] + 13
			}
		} else if arr[i] >= 'a' && arr[i] <= 'z' {
			if (arr[i] + 13) > 'z' {
				targetUnicodeNumber := 'a' + (arr[i] + 12 - 'z')
				arr[i] = targetUnicodeNumber
			} else {
				arr[i] = arr[i] + 13
			}
		}
	}
	return len(arr), io.EOF
}
