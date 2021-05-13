package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	callTest()
}

func callTest() {
	httpRequest, err := http.NewRequest(http.MethodGet, "http://localhost:8080/test", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	httpResponse, err := httpClient.Do(httpRequest)
	if httpResponse != nil && httpResponse.Body != nil {
		defer func() {
			_ = httpResponse.Body.Close()
		}()
	}

	if err != nil  {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			fmt.Println("error timeout ")
			return
		}
		fmt.Println("others error")
		return
	}

	if httpResponse.StatusCode == http.StatusRequestTimeout {
		fmt.Println("timeout")
	}

	bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bodyBytes))

}