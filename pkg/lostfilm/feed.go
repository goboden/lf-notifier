package lostfilm

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func fetchXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return make([]byte, 0), err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}

func parseXML(body []byte) RSS {
	var rss RSS
	xml.Unmarshal(body, &rss)

	return rss
}

func Subscribe(url string, interval uint32) chan Item {
	channel := make(chan Item)

	go func() {
		for {
			time.Sleep(time.Second * time.Duration(interval))

			body, err := fetchXML(url)
			if err != nil {
				log.Printf("Fetch error: %s", err.Error())
			}

			rss := parseXML(body)
			for _, item := range rss.Channel.Items {
				channel <- item
			}
		}
	}()

	return channel
}

func read(url string, interval uint32, data chan Item) {

}
