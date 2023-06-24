package reader

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nikonor/quickgobot/writer"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r Reader) Handler(idx int, wg *sync.WaitGroup, doneChan <-chan struct{}, updates <-chan tgbotapi.Update,
	wChan chan writer.Message) {
	for {
		select {
		case <-doneChan:
			wg.Done()
			return
		case u := <-updates:
			wChan <- writer.Message{
				ChatId:  u.Message.Chat.ID,
				Message: u.Message.Text,
			}
		}
	}
}

func fill() string {
	return `   A B C D E F G H J K L M N O P Q R S T
19 . . . . . . . . . . . . . . . . . . . 19
18 . . . . . . . . . . . . . . . . . . . 18
17 . . . . . . . . . . . . . . . . . . . 17
16 . . . + . . . . . + . . . . . + . . . 16
15 . . . . . . . . . . . . . . . . . . . 15
14 . . . . . . . . . . . . . . . . . . . 14
13 . . . . . . . . . . . . . . . . . . . 13
12 . . . . . . . . . . . . . . . . . . . 12
11 . . . . . . . . . . . . . . . . . . . 11
10 . . . + . . . . . + . . . . . + . . . 10
 9 . . . . . . . . . . . . . . . . . . . 9
 8 . . . . . . . . . . . . . . . . . . . 8
 7 . . . . . . . . . . . . . . . . . . . 7
 6 . . . . . . . . . . . . . . . . . . . 6
 5 . . . . . . . . . . . . . . . . . . . 5
 4 . . . + . . . . . + . . . . . + . . . 4
 3 . . . . . . . . . . . . . . . . . . . 3
 2 . . . . . . . . . . . . . . . . . . . 2
 1 . . . . . . . . . . . . . . . . . . . 1
   A B C D E F G H J K L M N O P Q R S T `
}
func fill2() string {
	return `
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .
. . . + . . . . . + . . . . . + . . .
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .
. . . + . . . . . + . . . . . + . . .
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .
. . . + . . . . . + . . . . . + . . .
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . .`
}
