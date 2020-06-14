package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const helpTxt = `Usage of %s: 
gtg <cmd> [options] <content>
cmd: 
	md5
	sha256
	base64
	jpg2png
	png2jpg
	uuid

options:
	-f file path
	-u result convert to upper case
	-h help

example1:
	gtg md5 123456
	output: d41d8cd98f00b204e9800998ecf8427e

example2:
	gtg md5 -u -f /home/jacpy/test.txt
	output: CB23B4F97510A693BA24CCC1ECC9A5F5
`

var (
	f bool
	u bool
	h bool
)

func processComm() bool {
	flag.BoolVar(&f, "f", false, "file path")
	flag.BoolVar(&u, "u", false, "upper case")
	flag.BoolVar(&h, "h", false, "help")

	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), helpTxt, os.Args[0])
	}

	// parse from third args, replace flag.Parse()
	if err := flag.CommandLine.Parse(os.Args[2:]); err != nil {
		fmt.Println(err)
		return true
	}

	if h {
		flag.Usage()
		return true
	}

	return false
}

func readFilePart(file *os.File, callback func([]byte)) error {
	buf := make([]byte, 8192)
	for {
		if n, err := file.Read(buf); err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println(err)
			return err
		} else {
			callback(buf[:n])
		}
	}

	return nil
}

func readFile(filePath string, write func(buf []byte), finish func() []byte) (string, error) {
	if file, err := os.Open(filePath); err != nil {
		log.Println(err)
		return "", err
	} else if err = readFilePart(file, func(bytes []byte) {
		write(bytes)
	}); err != nil {
		log.Println(err)
		return "", err
	} else {
		result := finish()
		return hex.EncodeToString(result), nil
	}
}
