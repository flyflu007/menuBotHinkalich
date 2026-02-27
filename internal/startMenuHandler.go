package internal

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"tg_bot/internal/keyboards"
)

func StartMenuHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Я - твой помощник для успешной работы в старике Хинкалыче",
		ParseMode:   models.ParseModeMarkdownV1,
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.MenuKeyboard()},
	})

}
