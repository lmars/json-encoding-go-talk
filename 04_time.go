package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println("Before:", t1)

	data, err := json.Marshal(t1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON:  ", string(data))

	var t2 time.Time
	if err := json.Unmarshal(data, &t2); err != nil {
		log.Fatal(err)
	}
	fmt.Println("After: ", t2)
}
