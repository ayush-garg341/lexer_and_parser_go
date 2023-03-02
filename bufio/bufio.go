package bufio

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Printf("Writing : %s\n", p)
	return len(p), nil
}

func BufferWriter() {
	// declare a buffered writer
	// with size 4
	w := new(Writer)
	bw := bufio.NewWriterSize(w, 4)

	// Case 1: Writing to buffer until full
	bw.Write([]byte{'1'})
	bw.Write([]byte{'2'})
	bw.Write([]byte{'3'})
	bw.Write([]byte{'4'}) // Write to io, buffer is full

	// Case 2: Buffer has space (4 bytes), as we already wrote to io
	bw.Write([]byte{'5'})
	err := bw.Flush() // Forcefully writing ( 1 byte )
	if err != nil {
		panic(err)
	}

	// Case 3: large write for buffer
	// Skip buffer and write directly to io
	bw.Write([]byte("12345"))
	available := bw.Available() // Should be 4 as we are directly writing to io
	fmt.Printf("Available space in buffer : %d\n", available)

	// We can re-use the same bufio.NewWriterSize for different writers using the reset() method:
	// writerOne := new(Writer)
	// bw := bufio.NewWriterSize(writerOne,2)
	// writerTwo := new(Writer)
	// bw.Reset(writerTwo)
}

func BufferReader() {
	// bufio allows us to read in batches with bufio.Reader
	// After a read data is released from buffer
	const singleLine string = "I'd love to have some tea right about now"
	const multiLine string = "Reading is my...\r\n favourite"
	fmt.Println("Length of single line input is " + strconv.Itoa(len(singleLine)))

	str := strings.NewReader(singleLine)
	br := bufio.NewReaderSize(str, 25)

	fmt.Println("\n---Peek---")

	// Peek - Case 1: Simple peek implementation
	b, err := br.Peek(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", b)

	// Peek - Case 2: Peek larger than buffer size
	b, err = br.Peek(30)
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}

	// Peek - Case 3: Buffer size larger than string
	br_large := bufio.NewReaderSize(str, 50)
	b, err = br_large.Peek(50)
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}

	// ReadSlice
	fmt.Println("\n---ReadSlice---")
	str = strings.NewReader(multiLine)
	r := bufio.NewReader(str)
	for {
		token, err := r.ReadSlice('.')
		if len(token) > 0 {
			fmt.Printf("Token (ReadSlice): %q\n", token)
		}
		if err != nil {
			fmt.Println("ReadSlice Error : " + err.Error())
			break
		}
	}

	// ReadLine
	fmt.Println("\n---ReadLine---")
	str = strings.NewReader(multiLine)
	r = bufio.NewReader(str)
	for {
		token, isPrefix, err := r.ReadLine()
		fmt.Println("Readline isPrefix : ", isPrefix)
		if len(token) > 0 {
			fmt.Printf("Token (ReadLine): %q\n", token)
		}
		if err != nil {
			fmt.Println("ReadLine Error : " + err.Error())
			break
		}
	}

	// ReadByte
	fmt.Println("\n---ReadByte---")
	str = strings.NewReader(multiLine)
	r.Reset(str)
	for {
		token, err := r.ReadBytes('\n')
		fmt.Printf("Token (ReadBytes): %q\n", token)
		if err != nil {
			fmt.Println("ReadBytes Error : " + err.Error())
			break
		}
	}

	// Scanner
	fmt.Println("\n---Scanner---")
	// Scanning stops at EOF, at first IO error, or if a token is too large to fit into the buffer
	str = strings.NewReader(multiLine)
	scanner := bufio.NewScanner(str)
	for scanner.Scan() {
		fmt.Printf("Token (Scanner): %q\n", scanner.Text())
	}
}
