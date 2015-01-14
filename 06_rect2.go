package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Rect struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func main() {
	rect := Rect{10, 20}
	fmt.Printf("decoded rect: %+v\n", rect)
	data, err := json.Marshal(rect)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("encoded rect:", string(data))

	fmt.Println("---")

	data = []byte(`{"x": 5, "y": 8}`)
	fmt.Println("encoded rect:", string(data))
	if err := json.Unmarshal(data, &rect); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decoded rect: %+v\n", rect)
}
