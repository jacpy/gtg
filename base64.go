package main

import (
	"encoding/base32"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func Base64(s string) error {
	err := base(s, func(buf []byte) {
		result := base64.StdEncoding.EncodeToString(buf)
		fmt.Println(result)
	}, func(s string) error {
		if result, err := base64.StdEncoding.DecodeString(s); err != nil {
			log.Println(err)
			return err
		} else {
			fmt.Println(string(result))
			return nil
		}
	})

	return err
}

func Base32(s string) error {
	err := base(s, func(buf []byte) {
		result := base32.StdEncoding.EncodeToString(buf)
		fmt.Println(result)
	}, func(s string) error {
		if result, err := base32.StdEncoding.DecodeString(s); err != nil {
			log.Println(err)
			return err
		} else {
			fmt.Println(string(result))
			return nil
		}
	})
	return err
}

func base(s string, encode func(buf []byte), decode func(s string) error) error {
	var (
		e bool
		d bool
	)

	flag.BoolVar(&e, "e", false, "encode")
	flag.BoolVar(&d, "d", false, "decode")

	if processComm() {
		return nil
	}

	if !e && !d || e && d {
		return errors.New("unsupported operation, encode and decode must not be same")
	}

	if f {
		if e {
			if buf, err := ioutil.ReadFile(s); err != nil {
				log.Println(err)
				return err
			} else {
				encode(buf)
				return nil
			}
		}

		if d {
			return errors.New("unsupported operation decode file")
		}
	} else {
		if e {
			encode([]byte(s))
		} else if d {
			return decode(s)
		}
	}

	return nil
}
