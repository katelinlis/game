package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/katelinlis/BackendMasters/internal"
)

func main() {

	go func() {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		opts := []bot.Option{
			bot.WithDefaultHandler(handler),
		}

		b, err := bot.New("", opts...)
		if err != nil {
			panic(err)
		}

		b.Start(ctx)
	}()

	server := internal.Init()
	server.Start()
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil && update.Message.Text == "/start" {

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   update.Message.Text,
			ReplyMarkup: models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{
						models.InlineKeyboardButton{Text: "открыть", WebApp: &models.WebAppInfo{URL: "https://game.katelinlis.com"}},
					},
				},
			},
		})
	}
}
