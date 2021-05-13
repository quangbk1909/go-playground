package main

import (
  "github.com/gin-gonic/gin"
  "test-pakage-go/pkga"
)

func main() {
  _ = gin.New()

  a:= pkga.NewUnexportA()
  a.



}
