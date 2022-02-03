package main

import (
	"fmt"
	"time"

	"github.com/goboden/lf-notifier/pkg/lostfilm"
)

func main() {
	dataChan := lostfilm.Subscribe("https://www.lostfilm.tv/rss.xml", 2)
	fmt.Printf("%v Subscribed\n", time.Now().Format("2006/01/02 15:04:05"))
	for {
		item := <-dataChan
		fmt.Printf("Item: %s", item.Title)
	}
}
