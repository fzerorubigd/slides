package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

// main START OMIT
// f START OMIT
func main() {
	id := make(chan string)

	go func() {
		h := sha1.New()
		c := []byte(time.Now().String())
		for {
			h.Write(c)
			id <- fmt.Sprintf("%x", h.Sum(nil))
		}
	}()
	fmt.Printf("This program will get 10 IDs from the id channel, print them and terminate\n\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", <-id)
	}
	
}
// f END OMIT
// main END OMIT
