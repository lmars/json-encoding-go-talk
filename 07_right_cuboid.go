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
		Rect{10, 5}, Rect{5, 7}, Rect{7, 10},
	}
	data, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
