package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	values := []interface{}{
		"hello",
		42,
		true,
		false,
		nil,
		map[string]interface{}{"foo": 42, "bar": true},
		[]interface{}{42, "string", false},
	}

	fmt.Println("Marshal")
	fmt.Println("-------")
	fmt.Println()

	for _, v := range values {
		encoded, err := json.Marshal(v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v => %s\n\n", v, string(encoded))
	}
}
