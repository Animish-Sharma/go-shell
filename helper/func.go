package helper

import (
	"fmt"
	"io"
	"os"
)

func CopyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func CopyStdin() {
	buf := make([]byte, 4096)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Fprintf(os.Stderr, "Err: Could not read the file %v\n", err)
				os.Exit(1)
			}
		}
		fmt.Print(string(buf[:n]))
	}
}
