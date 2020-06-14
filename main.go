package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Println(helpTxt)
		return
	}

	content := os.Args[len(os.Args)-1]
	var err error = nil
	cmdName := strings.ToLower(os.Args[1])
	switch cmdName {
	case "md5":
		err = Md5(content)
	case "sha256":
		err = Sha256(content)
	case "sha384":
		err = sha384(content)
	case "sha512":
		err = sha512(content)
	case "base64":
		err = Base64(content)
	case "uuid":
		err = MakeUuid()
	case "jpg2png":
		err = jpg2png(content)
	case "png2jpg":
		err = png2jpg(content)
	default:
		err = errors.New("unsupported command " + cmdName)
	}

	if err != nil {
		log.Println(err)
	}
}
