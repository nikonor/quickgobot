package main

import (
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nikonor/quickgobot/conf"
	"github.com/nikonor/quickgobot/reader"
	"github.com/nikonor/quickgobot/writer"
)

type Reader interface {
	Handler(idx int, wg *sync.WaitGroup, doneChan <-chan struct{}, updates <-chan tgbotapi.Update, wChan chan writer.Message)
}

type Writer interface {
	Handler(idx int, wg *sync.WaitGroup, doneChan <-chan struct{}, tbot *tgbotapi.BotAPI)
}

func main() {
	token, ok := os.LookupEnv("TLG_TOKEN")
	if !ok {
		panic("wrong token")
	}

	cfg, err := conf.Load("./config.json")
	if err != nil {
		panic(err.Error())
	}

	tbot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err.Error())
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// TODO: webhook

	whInfo, err := tbot.GetWebhookInfo()
	if err != nil {
		panic(err.Error())
	}

	if whInfo.IsSet() {
		if _, err = tbot.Send(tgbotapi.DeleteWebhookConfig{DropPendingUpdates: true}); err != nil {
			panic(err.Error())
		}
	}
	doneChan := make(chan struct{})
	updates := tbot.GetUpdatesChan(u)
	wg := new(sync.WaitGroup)
	r := reader.NewReader()
	w := writer.NewWriter()
	ch := w.Chan()

	// TODO: слушаем сигналы

	for i := 0; i < cfg.Workers; i++ {
		wg.Add(2)
		go w.Handler(i+1, wg, doneChan, tbot)
		go r.Handler(i+1, wg, doneChan, updates, ch)
	}
	wg.Wait()
}
