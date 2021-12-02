package go_telegram_rss_bot

import (
	"fmt"
	"time"

	"github.com/anaskhan96/soup"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mmcdole/gofeed/rss"
)

type ParseFunc func(*FeedPair, *rss.Item, int64)

var parse_mode = map[string]ParseFunc{
	"cnBeta":    Parse_cnBeta,
	"economist": Parse_economist,
}

func RegisterParseMode(name string, callback ParseFunc) {
	parse_mode[name] = callback
}

func (feed_pair *FeedPair) Parse(items []*rss.Item, chat_id int64) {
	feeds_len := len(items)
	if feeds_len > cfg.FeedMaxLimit {
		feeds_len = cfg.FeedMaxLimit
	}
	items = items[:feeds_len]

	this_parse_mode := parse_mode[feed_pair.Name]
	if this_parse_mode == nil {
		this_parse_mode = Parse_Normal
	}
	for i := range items {
		if feed_pair.IsLast(items[i].Link) {
			break
		}
		fmt.Printf("%d - %s\n", i+1, items[i].Title)
		this_parse_mode(feed_pair, items[i], chat_id)
		time.Sleep(2 * time.Second)
	}
	feed_pair.SetLast(items[0].Link)
}

// Add your parse mode here.
func Parse_cnBeta(feed_pair *FeedPair, item *rss.Item, chat_id int64) {
	title := item.Title
	link := item.Link

	text := fmt.Sprintf("*[%s]*\n[%s](%s)\n", feed_pair.Name, title, link)
	msg := tgbot.NewMessage(chat_id, text)
	msg.ParseMode = "Markdown"
	SendMessage(msg)
	// fmt.Println(text)
}

func Parse_economist(feed_pair *FeedPair, item *rss.Item, chat_id int64) {
	content := item.Description
	title := item.Title
	link := item.Link

	text := fmt.Sprintf("*[%s]*\n[%s](%s)\n\n`%s`\n", feed_pair.Name, title, link, content)
	msg := tgbot.NewMessage(chat_id, text)
	msg.ParseMode = "Markdown"
	SendMessage(msg)
	// fmt.Println(text)
}

func Parse_Normal(feed_pair *FeedPair, item *rss.Item, chat_id int64) {
	content := item.Description
	title := item.Title
	link := item.Link
	rss_content := soup.HTMLParse(content)

	// if the description has no <p>, this will cause a crash
	first_paragraph := rss_content.Find("p").FullText()
	index := len(first_paragraph)
	for i := range first_paragraph {
		if first_paragraph[i] != '\n' && first_paragraph[i] != ' ' {
			index = i
			break
		}
	}

	text := fmt.Sprintf("*[%s]*\n[%s](%s)\n\n`%s`\n", feed_pair.Name, title, link, first_paragraph[index:])
	msg := tgbot.NewMessage(chat_id, text)
	msg.ParseMode = "Markdown"
	SendMessage(msg)
	// fmt.Println(text)
}
