package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

func LineCounter(f string) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	file, err := os.Open(f)
	if err != nil {
		log.Fatal("error: ", err)
	}

	fileSatae, err := file.Stat()
	if err != nil {
		log.Fatal("error: ", err)
	}

	if fileSatae.Size() == 0 {
		return -1, fmt.
	}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

/*
func LineCounter(r io.Reader) (int, error) {
	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}
*/
