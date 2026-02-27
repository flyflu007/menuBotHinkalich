package internal

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"tg_bot/internal/keyboards"
)

func MainMainHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
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

func MainDescriptionHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            "",
	})
	var PhotoURL string
	var CaptionDelish string
	switch update.CallbackQuery.Data[5:] {
	case "plov":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/cff5dbd552ea3b5c057fbb5ee7b5d0cd.PNG"
		CaptionDelish = "*Плов с говядиной 300 г*\n" +
			"Рассыпчатый ароматный рис с говядиной, овощами,\n" +
			" специями и барбарисом\n" +
			"*Украшение:* перец чили консервированный\n\n" +
			"*Состав:* Рис, барбарис, говядина мякоть, лук репчатый,\n" +
			"морковь, чеснок, зира, соль, куркума, масло растительное."
	case "chahoh":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/18d2ecd8ba5daa9c421b1fa70f48e595.PNG"
		CaptionDelish = "*Чахохбили из курицы 300 г*\n" +
			"Курица, тушеная в томатном пряном соусе с добавлением овощей,\n" +
			"чеснока, аджики,грецкого ореха, ароматных грузинских специй и свежей зелени.\n" +
			"*Украшение:* кинза, перец чили маринованный\n\n" +
			"*Состав:* Куриное филе бедра, лук репчатый, помидор,\n" +
			"перец болгарский, орех грецкий, томатная паста,\n" +
			"чеснок, аджика, масло подсолнечное, кориандр\n," +
			"уцхо-сунели, хмели-сунели, сахар, соль, перец,\n" +
			"зелень(базилик, петрушка, кинза"
	case "minihin":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/d4c128d5ae9fb824e117583fa663cbef.PNG"
		CaptionDelish = "*Мини-хинкали с курицей в сурном соусе 290 г*\n" +
			"10 мини-хинкали с курицей в нежном сливочно-сырном соусе.\n" +
			"*Украшение:* кинза\n\n" +
			"Тесто на хинкали (мука, вода, соль), фарш(куриный фарш, сливки, соль),\n" +
			"сырный соус(сыр плавленый, сливки растительное, вода питьевая,\n" +
			"масло сливочное, соль, хмели-сунели."
	case "odjahur":
		PhotoURL = "https://github.com/flyflu007/menuBotHinkalich/blob/main/odjahuri.png"
		CaptionDelish = "*Оджахури со свининой 300 г*\n" +
			"Обжаренный картофель со свининой и луком, ароматными специями,\n" +
			"чесноком и укропом." +
			"*Украшение:* укроп\n\n" +
			"*Состав:* Картофель отварной в мундире, свинина жареная с луком,\n" +
			"масло подсолнечное, соль, хмели-сунели, перец, черный молотый, чеснок,\n" +
			"укроп."

	case "lagman":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/c9f0872c34f4fffac644dd0044aec9a4.PNG"
		CaptionDelish = "*Лагман с говядиной 300 г*\n" +
			"Густое блюдо с говядиной, лапшой и овощами.\n" +
			"*Украшение:* нарезанная зелень укропа и маринованный перец.\n\n" +
			"*Состав:* Говядина, перец болгарский, баклажан, морковь, картофель, лук, чеснок,\n" +
			"томатная паста, соевый соус, соль, сахар, специи(бадьян, перец черный, зира, лавровый лист),\n" +
			"масло растительное, говяжий бульон, лапша."
	case "dolma":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/9771250e946310187ff6b9b8b6d62a8c.PNG"
		CaptionDelish = "*Долма 220 г*\n" +
			"Виноградные листья, начиненные мясным фаршем и рисом с добавлением аджики и ароматных специй,\n" +
			"В порции 7-12 штук(блюдо подается по весу) подается с мацони.\n\n" +
			"*Состав:* Свино-говяжий фарш, аджика, рис, соль, перец черный, чеснок\n" +
			"лук, тмин, кориандр, уцхо-сунели, мацони(молоко, закваска)."
	case "stake":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/f2095ab41cb60c2cfd76f3dfe8d6ceda.PNG"
		CaptionDelish = "*Стейк из капусты с ореховым соусом 230 г*\n" +
			"Ломтик белокачанной купусты, запечённый в горчичном соусе. Подается с ореховым соусом и зеленым маслом.\n" +
			"Украшается зеоёным маслом и веточкоц темьяна.\n" +
			"*Украшение:* масло зеленое, веточка тимьяна\n\n" +
			"*Состав:* Капуста белокачанная, соль, соус Табуле, копченая паприка,\n" +
			"соус ореховый(соус Баже, вода питьевая, соль)."
	}
	b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		Photo:     &models.InputFileString{Data: PhotoURL},
		Caption:   CaptionDelish,
		ParseMode: models.ParseModeMarkdownV1,
	})
}
