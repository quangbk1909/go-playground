package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/guregu/null.v4"
)

func main() {
	var slice []string
	fmt.Println(slice == nil)

	test := test{
		PrimitiveType: "",
		Slice:        slice,
	}

	testJson, _ := json.Marshal(test)
	fmt.Println(string(testJson))
	b := B{}
//	jsonString := `{
//	"field": ""
//}`
//	jsonByte := []byte(jsonString)
//	err := json.Unmarshal(jsonByte, &b)
//	if err != nil {
//		panic(err)
//	}
	bByte, _ := json.Marshal(b)
	fmt.Println(string(bByte))
}

type test struct {
	PrimitiveType string   `json:"primitive_type,omitempty"`
	Slice         []string `json:"slice,omitempty"`
	StructField   A        `json:"struct_field,omitempty"`
}

type A struct {
	Field string `json:"field,omitempty"`
}

type B struct {
	Field null.String `json:"field,omitempty"`
}
