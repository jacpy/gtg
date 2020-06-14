package main

import (
	"crypto/md5"
	"crypto/sha256"
	sha5122 "crypto/sha512"
	"encoding/hex"
	"hash"
	"log"
	"strings"
)

func Md5(s string) error {
	return hashCode(md5.New(), s)
}

func Sha256(s string) error {
	return hashCode(sha256.New(), s)
}

func sha384(s string) error {
	return hashCode(sha5122.New384(), s)
}

func sha512(s string) error {
	return hashCode(sha5122.New(), s)
}

func hashCode(target hash.Hash, s string) error {
	processComm()

	var hexStr string
	if f {
		if str, err := readFile(s, func(buf []byte) {
			target.Write(buf)
		}, func() []byte {
			return target.Sum(nil)
		}); err != nil {
			log.Println(err)
			return err
		} else {
			hexStr = str
		}
	} else {
		target.Write([]byte(s))
		result := target.Sum(nil)
		hexStr = hex.EncodeToString(result)
	}

	if u {
		hexStr = strings.ToUpper(hexStr)
	}

	log.Println(hexStr)
	return nil
}
