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

	str := t1.Format(time.RFC3339Nano)
	data, err := json.Marshal(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON:  ", string(data))

	if err := json.Unmarshal(data, &str); err != nil {
		log.Fatal(err)
	}
	t2, err := time.Parse(time.RFC3339Nano, str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After: ", t2)
}
