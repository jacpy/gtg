package main

import (
	"errors"
	"flag"
)

func Base64(s string) error {
	var (
		e bool
		d bool
	)

	flag.BoolVar(&e, "e", false, "encode")
	flag.BoolVar(&d, "d", false, "decode")

	if !e && !d || e && d {
		return errors.New("unsupported operation, encode and decode must not be same")
	}

	return nil
}
