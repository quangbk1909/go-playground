package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("http://localhost:8080/webhook/test")
	if err != nil {
		panic(err)
	}
	
	fmt.Println(ioutil.ReadAll(resp.Body))
}


