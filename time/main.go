package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second)
	end := time.Now()
	d := end.Sub(start)
	fmt.Printf("%d",d)
}
