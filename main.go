package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func convert(input []byte) []byte {
	output := bytes.Buffer{}
	for len(input) > 0 {
		if len(input) < 9 {
			output.Write(input)
			break
		}
		segment := input[:9]
		if string(segment) == "_binary '" {
			input = input[9:]
			output.Write([]byte("0x"))
			for len(input) > 0 {
				if len(input) >= 2 {
					s := string(input[:2])
					if s == "'," {
						input = input[2:]
						output.Write([]byte(","))
						break
					}
					if s == "')" {
						input = input[2:]
						output.Write([]byte(")"))
						break
					}
					if s == "\\0" {
						output.Write([]byte("00"))
						input = input[2:]
						continue
					}
				}

				output.Write([]byte(fmt.Sprintf("%02X", input[0])))
				input = input[1:]
			}
			continue
		}
		output.Write([]byte{input[0]})
		input = input[1:]
	}

	return output.Bytes()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// output to stdio
	_, _ = io.Copy(os.Stdout, bytes.NewReader(convert(data)))
}
