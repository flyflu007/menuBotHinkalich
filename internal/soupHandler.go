package internal

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"tg_bot/internal/keyboards"
)

func SoupMainHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.Message.ID,
		Text:        "Список всех супов",
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.SoupKeyboard()},
	})

}

func SoupDescriptionHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            "",
	})
	var PhotoURL string
	var CaptionDelish string
	switch update.CallbackQuery.Data[5:] {
	case "harcho":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/686b796ff6459d8ee56115662ac9973e.PNG"
		CaptionDelish = "*Харчо с говядиной 250 г*\n" +
			"Густой пряный суп на основе говяжего бульона\n" +
			" и томата с отварной говядиной, рисом, грецкими орехами, чесноком,\n" +
			"свежей зеленью и ароматными специями.\n" +
			"*Украшение:* кинза.\n\n" +
			"*Состав:* овядина, рис белый круглый, уцхо-сунели,\n" +
			" хмели - сунели, кориандр, чеснок, лук репчатый, кинза,\n" +
			" соль, сахар, перец, орех грецкий, аджика, масло подсолнечное,\n" +
			" масло сливочное, зафаран, томатная паста, паприка."
	case "cremsoup":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/af631e8eabe10b18408c8632f4f0636c.PNG"
		CaptionDelish = "*Крем-суп чкмерули 250 г* \n" +
			"Ароматный сливочно-чесночный суп с добавлением курицы и специй\n" +
			"*Украшение:* оранжевое масло, темьян\n\n" +
			"Сливки растительные, молоко, бедро куриное, мука,\n" +
			" филе куриное, морковь.\n" +
			"Специи: уцхо-сунели, хмели-сунели, кориандр, тимьян,\n" +
			" соль поваренная, чеснок."
	case "chihir":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/a15b591103c4b312468655c43f50d32b.PNG"
		CaptionDelish = "*Чихиртма с курицей 250 г*\n" +
			"Наваристый куриный суп с куриной мякотью, загущенным яйцом\n" +
			"и свежей ароматной зеленью.\n" +
			"*Украшение:* мята\n\n" +
			"*Состав:* Мясо курицы, яйцо кур, лук репчатый,\n" +
			" кинза, петрушка, сельдерей, масло растительное,\n" +
			" масло сливочное, зафаран, уксус винный, перец, соль."
	case "grib":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/0c2ba684979841e5791ceb4f0cc276fb.PNG"
		CaptionDelish = "*Грибная южка 250 г*\n" +
			"Насыщенный суп на основе двух видов грибов с морковью\n" +
			"и зеленью.\n" +
			"*Украшение:* листок петрушки\n\n" +
			"*Состав:* Грибы шампиньоны, грибы вешанки, лук репчатый,\n" +
			" морковь, масло подсолнечное, мука, соль, перец черный, лавровый лист,\n" +
			" грибная приправа, укроп, петрушка."
	}
	b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		Photo:     &models.InputFileString{Data: PhotoURL},
		Caption:   CaptionDelish,
		ParseMode: models.ParseModeMarkdownV1,
	})

}
