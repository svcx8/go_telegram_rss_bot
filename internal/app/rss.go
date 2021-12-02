package go_telegram_rss_bot

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mmcdole/gofeed/rss"
)

type ContentParser interface {
	Parse(*FeedPair, []*rss.Item)
}

func StartFetchRss() {
	main_func := func(feed_list []FeedPair, interval int, chat_id int64) {
		for {
			for _, feed := range feed_list {
				if res, err := http.Get(feed.Link); err == nil {
					if res.StatusCode == 200 {
						rss_parser := rss.Parser{}
						rss_items, _ := rss_parser.Parse(res.Body)
						feed.Parse(rss_items.Items, chat_id)
					}
					res.Body.Close()
				}
			}
			time.Sleep(time.Duration(interval) * time.Minute)
		}
	}
	fast_chat_id, _ := strconv.ParseInt(cfg.TgFastChannel, 10, 64)
	slow_chat_id, _ := strconv.ParseInt(cfg.TgSlowChannel, 10, 64)

	if len(sub.FastList) > 0 {
		go main_func(sub.FastList, cfg.FastInterval, fast_chat_id)
	}

	if len(sub.SlowList) > 0 {
		go main_func(sub.SlowList, cfg.SlowInterval, slow_chat_id)
	}
}

func ParseRssTest() {
	file_name := "raw.xml"
	fd, err := os.Open(file_name)
	if err != nil {
		log.Panicf("Cannot open file %s.", file_name)
	}

	rss_parser := rss.Parser{}
	rss_feed, _ := rss_parser.Parse(fd)

	feed_pair := FeedPair{Name: "economist"}
	feed_pair.Parse(rss_feed.Items, int64(cfg.TgAdminId))
}
