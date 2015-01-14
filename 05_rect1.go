package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Rect struct {
	X int
	Y int
}

func (r *Rect) MarshalJSON() ([]byte, error) {
	lengths := map[string]int{
		"x": r.X,
		"y": r.Y,
	}
	return json.Marshal(lengths)
}

func (r *Rect) UnmarshalJSON(data []byte) error {
	var lengths map[string]int
	if err := json.Unmarshal(data, &lengths); err != nil {
		return err
	}
	r.X = lengths["x"]
	r.Y = lengths["y"]
	return nil
}

func main() {
	rect := Rect{10, 20}
	fmt.Printf("decoded rect: %+v\n", rect)
	data, err := json.Marshal(&rect)
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
