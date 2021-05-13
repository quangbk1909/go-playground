package main

import (
	"fmt"
	"github.com/h2non/filetype"
	"os"
)

func main() {
	file, err := os.Open("test-odt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err = file.Read(buff)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	kind, _ := filetype.Match(buff)
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
		return
	}

	fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)

	//buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	//_, err = file.Read(buff)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//filetype := http.DetectContentType(buff)
	//
	//fmt.Println(filetype)

	//switch filetype {
	//case "image/jpeg", "image/jpg":
	//	fmt.Println(filetype)
	//
	//case "image/gif":
	//	fmt.Println(filetype)
	//
	//case "image/png":
	//	fmt.Println(filetype)
	//
	//case "application/pdf":       // not image, but application !
	//	fmt.Println(filetype)
	//default:
	//	fmt.Println("unknown file type uploaded")
	//}
}
