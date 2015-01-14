package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	m := map[string]interface{}{"foo": 42, "bar": true}
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(m); err != nil {
		log.Fatal(err)
	}

	rd, wr, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(rd)
	go func() {
		// slowly write a JSON array of ints to wr
		for i := 0; i < 10; i++ {
			fmt.Fprintf(wr, "[%d]", i)
			time.Sleep(1000 * time.Millisecond)
		}
		wr.Close()
	}()

	for {
		var arr []int
		if err := dec.Decode(&arr); err != nil {
			break
		}
		fmt.Printf("%+v\n", arr)
	}
}
