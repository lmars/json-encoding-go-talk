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

type RightCuboid struct {
	A Rect `json:"a"`
	B Rect `json:"b"`
	C Rect `json:"c"`
}

func main() {
	r := RightCuboid{
		Rect{2, 3}, Rect{4, 5}, Rect{6, 7},
	}
	data, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
