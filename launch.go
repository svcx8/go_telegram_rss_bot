package main

import (
	"fmt"
	rssbot "go_telegram_rss_bot/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	SetupCloseHandler()
	rssbot.LoadConfig()
	rssbot.ReadHistory()
	go rssbot.StartBot()
	select {}
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Good bye!")
		rssbot.SaveHistory()
		os.Exit(0)
	}()
}
