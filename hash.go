package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Md5(s string) string {
	md := md5.New()
	md.Write([]byte(s))
	result := md5.Sum(nil)
	return hex.EncodeToString(result[:])
}

func Sha256(s string) string {
	sh := sha256.New()
	sh.Write([]byte(s))
	result := sh.Sum(nil)
	h := hex.EncodeToString(result)
	return h
}
