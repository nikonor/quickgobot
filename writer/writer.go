package writer

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Writer struct {
	ch chan Message
}

func NewWriter() *Writer {
	return &Writer{
		ch: make(chan Message),
	}
}

type Message struct {
	ChatId  int64
	Message string
}

func (w Writer) Handler(idx int, wg *sync.WaitGroup, doneChan <-chan struct{}, tbot *tgbotapi.BotAPI) {
	for {
		select {
		case <-doneChan:
			return
		case msg := <-w.ch:
			m := tgbotapi.NewMessage(msg.ChatId, "<code>"+msg.Message+"</code>")
			m.ParseMode = tgbotapi.ModeHTML
			tbot.Send(m)
		}
	}
}

func (w Writer) Chan() chan Message {
	return w.ch
}
