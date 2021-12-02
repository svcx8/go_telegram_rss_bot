package go_telegram_rss_bot

import (
	"encoding/json"
	"log"
	"os"
)

var rss_history map[string]string

func (feed_pair *FeedPair) SetLast(url string) {
	rss_history[feed_pair.Name] = url
}

func (feed_pair *FeedPair) IsLast(url string) bool {
	return rss_history[feed_pair.Name] == url
}

func SaveHistory() {
	res, _ := json.Marshal(&rss_history)
	os.WriteFile("history.json", res, 0755)
}

func ReadHistory() {
	fd, err := os.ReadFile("history.json") // if you're using for the first time, keep history.json equal `{}`
	if err != nil {
		log.Panic(err)
	}

	if err := json.Unmarshal(fd, &rss_history); err != nil {
		log.Panic(err)
	}
}
