package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func MakeUuid() error {
	if id, err := uuid.NewV4(); err != nil {
		return err
	} else {
		s := id.String()
		fmt.Println(s)
	}

	return nil
}
