package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var str string
	var num int
	var bTrue bool
	var bFalse bool
	var null struct{}
	var obj map[string]interface{}
	var arr []interface{}

	values := map[string]interface{}{
		`"hello"`: str,
		`42`:      num,
		`true`:    bTrue,
		`false`:   bFalse,
		`null`:    null,
		`{ "foo": 42, "bar": true }`: obj,
		`[ 42, "string", false ]`:    arr,
	}

	fmt.Println("Unmarshal")
	fmt.Println("---------")
	fmt.Println()

	for data, v := range values {
		err := json.Unmarshal([]byte(data), &v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %+v\n\n", data, v)
	}
}
