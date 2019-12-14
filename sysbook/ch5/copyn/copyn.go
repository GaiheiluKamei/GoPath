package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func CopyNOffset(dst io.Writer, src io.ReadSeeker, offset, length int64) (int64, error) {
	if _, err := src.Seek(offset, io.SeekStart); err != nil {
		return 0, err
	}
	return io.CopyN(dst, src, length)
}

func main() {
	r := strings.NewReader("This is an example of CopyN with offset")
	for i, l, step := int64(0), int64(r.Len()), int64(5); i < l; i += step {
		CopyNOffset(os.Stdout, r, i, step)
		fmt.Println()
	}
}
