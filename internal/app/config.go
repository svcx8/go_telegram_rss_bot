package go_telegram_rss_bot

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type UserConfig struct {
	TgBotToken    string `json:"tg_bot_token"`
	TgAdminId     int    `json:"tg_admin_id"`
	TgSlowChannel string `json:"tg_slow_channel"`
	TgFastChannel string `json:"tg_fast_channel"`
	FeedMaxLimit  int    `json:"feed_max_limit"`
	FastInterval  int    `json:"fast_interval"`
	SlowInterval  int    `json:"slow_interval"`
}

type FeedPair struct {
	Link string
	Name string
}

type Feeds struct {
	FastList []FeedPair `json:"fast_list"`
	SlowList []FeedPair `json:"slow_list"`
}

var cfg = UserConfig{
	"",
	123456,
	"",
	"",
	10,
	10,
	360,
}

var sub = Feeds{}

func LoadConfig() {
	fd, err := os.ReadFile("cfg.json")
	if err != nil {
		log.Panic(err)
	}

	if err := json.Unmarshal(fd, &cfg); err != nil {
		log.Panic(err)
	}

	fd, err = os.ReadFile("feeds.json")

	if err != nil {
		log.Panic(err)
	}

	if err = json.Unmarshal(fd, &sub); err != nil {
		log.Panic(err)
	}
}

func WriteHistoryTest() {
	rss_history = make(map[string]string)

	sub.FastList[0].SetLast("baidu~")
	sub.FastList[1].SetLast("bing~")

	fast_list_history_0 := rss_history[sub.FastList[0].Name]

	fmt.Println(fast_list_history_0)

	fmt.Println(sub.FastList[0].IsLast("world~"))
	fmt.Println(sub.FastList[0].IsLast("hello"))
	SaveHistory()
}

func ReadHistoryTest() {
	ReadHistory()
	for i := range rss_history {
		fmt.Printf("%s - %s\n", i, rss_history[i])
	}
}

func LoadFeedsTest() {
	sub.FastList = append(sub.FastList, FeedPair{"http://baidu.com/rss", "baidu"})
	sub.FastList = append(sub.FastList, FeedPair{"http://bing.com/rss", "bing"})
	sub.SlowList = append(sub.SlowList, FeedPair{"http://google.com/rss", "google"})
	sub.SlowList = append(sub.SlowList, FeedPair{"http://twitter.com/rss", "twitter"})

	bytes, err := json.Marshal(&sub)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(bytes))

	// ReadHistoryTest()
	// WriteHistoryTest()
}
