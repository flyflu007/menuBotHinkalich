package internal

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"tg_bot/internal/keyboards"
)

func HinkaliMainHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.Message.ID,
		Text:        "Список всех основых блюд",
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.HinakliKeyboard()},
	})

}

func HinkaliDescriptionHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            "",
	})
	var PhotoURL string
	var CaptionDelish string
	switch update.CallbackQuery.Data[8:] {
	case "barzel":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/e2b199913723c0c801d9a554cdfd4ba3.PNG"
		CaptionDelish = "*Баранина и зелень 85 г*\n" +
			"Хинкали с начинкой из бараньего фарша с аджикой, кинзой\n" +
			" и пряным насыщенным бульоном\n\n" +
			"*Состав:* Тесто (мука, вода, соль), фарш (свинина 70% \n" +
			"(65%-свинина; 5%-сало, говядина 30%, свиное сало, аджика,\n" +
			" луково-чесночный соус, соль, перец черный гор, кинза, тмин)."
	case "grib":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/e51fdf6b487c2192db6c117bffdb36d8.PNG"
		CaptionDelish = "Грибы 85 г*\n" +
			"Хинкали с начинкой из обжаренных грибов\n" +
			" и лука с добавлением сливок и специй\n\n" +
			"*Состав:* Тесто (мука, вода, соль), фарш (шампиньоны, лук репчатый, масло подсолнечное,\n" +
			" укроп, сливки, перец черный, кориандр, молоко)."
	case "svingovzel":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/8bea1381b46708a92b637860bab1e40c.PNG"
		CaptionDelish = "*Свинина и говядина с зеленью 85 г*\n" +
			"Хинкали с начинкой из свино-говяжьего фарша с аджикой, кинзой\n" +
			" и пряным насыщенным бульоном\n\n" +
			"*Состав:* Тесто (мука, вода, соль), фарш (свинина 70% (65%-свинина\n" +
			" 5%-сало, говядина 30%, свиное сало, аджика, луково-чесночный соус,\n" +
			" соль, перец черный гор, кинза, тмин)."
	case "svingov":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/2461fbefc407cfe94428deaf3ffedfe9.PNG"
		CaptionDelish = "*Свинина и говядина без зелени 85 г*\n" +
			"Хинкали с начинкой из бараньего фарша с аджикой, \n" +
			" и пряным насыщенным бульоном\n\n" +
			"*Состав:* Тесто (мука, вода, соль), фарш (свинина 70% \n" +
			"(65%-свинина; 5%-сало, говядина 30%, свиное сало, аджика,\n" +
			" луково-чесночный соус, соль, перец черный гор, тмин)."
	case "cheese":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/b07aa86b646df74f98a88e4360dd2493.PNG"
		CaptionDelish = "*Три сыра 85 г*\n" +
			"Хинкали с начинкой из трех видов домашнего сыра\n" +
			" с добавление сливок и специй\n\n" +
			"*Состав:* Тесто (мука, вода, соль), \n" +
			"фарш (три вида сыра: адыгейский, брынза, сулугуни, сливки,\n" +
			" уцхо-сунели, кориандр, молоко)."

	case "vishtvor":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/ee37ab77540b166a1465ea1b3007f113.PNG"
		CaptionDelish = "*Вишня и творог 85 г*\n" +
			"Хинкали со сладкой начинкой из вишни и творога\n" +
			" *Осторожно: может содержать косточку* \n\n" +
			"*Состав:* Тесто (мука, вода, соль), фарш (вишня, сахар,\n" +
			"крахмал картофельный, творог)."

	case "gov":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/ce1a32ba94a9484e5cd37928ba4e3187.PNG"
		CaptionDelish = "*Говядина 85 г*\n" +
			"Тесто (мука, вода, соль), говядина, говяжий жир,\n" +
			" перец черный, соль, аджика, горчица по-дижонски,\n" +
			" уцхо-сунели, кориандр, паприка, чеснок, укроп."

	case "kur":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/93e859878dd1d2cc037321b27f00e327.PNG"
		CaptionDelish = "*Курица 85 г*\n" +
			"Хинкали с начиной их курицы, чеснока\n" +
			" и свежей петрушки \n\n" +
			"*Состав:* Тесто (мука, вода, соль),фарш (бедро куриное,\n" +
			" соль, перец чёрный молотый, зелень петрушки, чеснок очищенный)."

	case "ind":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/ef3528321aafa63d8be62e0e4f684d15.PNG"
		CaptionDelish = "*Индейка 85 г*\n" +
			"Хинкали с начиной их сочной индейки с зеленью\n\n" +
			"*Состав:* ???????Тесто (мука, вода, соль),фарш (индейка,\n" +
			" соль, перец чёрный молотый, чеснок очищенный).???????"

	}

	b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		Photo:     &models.InputFileString{Data: PhotoURL},
		Caption:   CaptionDelish,
		ParseMode: models.ParseModeMarkdownV1,
	})
}
