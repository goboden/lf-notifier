package lostfilm

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	Link        string   `xml:"link"`
}

func Load(ch chan string, quit chan bool) {
	url := "https://www.lostfilm.tv/rss.xml"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var rss RSS
	xml.Unmarshal(body, &rss)

	// fmt.Printf("%v", rss)
	for _, item := range rss.Channel.Items {
		ch <- item.Title
	}

	quit <- true
}
