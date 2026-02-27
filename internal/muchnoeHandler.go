package internal

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"tg_bot/internal/keyboards"
)

func MuchnoeMainHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.Message.ID,
		Text:        "Список всех основых блюд",
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.MainKeyboard()},
	})

}

func MuchnoeDescriptionHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            "",
	})
	var PhotoURL string
	var CaptionDelish string
	switch update.CallbackQuery.Data[7:] {

	}

	b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		Photo:     &models.InputFileString{Data: PhotoURL},
		Caption:   CaptionDelish,
		ParseMode: models.ParseModeMarkdownV1,
	})
}
