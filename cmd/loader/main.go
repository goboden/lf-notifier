package main

import (
	"fmt"

	"github.com/goboden/lf-notifier/pkg/lostfilm"
)

func main() {
	data := make(chan string, 1)
	quit := make(chan bool)

	go lostfilm.Load(data, quit)

	for {
		select {
		case msg := <-data:
			fmt.Println(msg)
		case q := <-quit:
			if q {
				return
			}
		}
	}
}
