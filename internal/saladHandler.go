package internal

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"tg_bot/internal/keyboards"
)

func SaladMainHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.Message.ID,
		Text:        "Список всех салатов",
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.SaladKeyboard()},
	})

}

func SaladDescriptionHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            "",
	})
	var PhotoURL string
	var CaptionDelish string
	switch update.CallbackQuery.Data[6:] {
	case "badrid":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/7c27cb36e7c1c60e19262d61a148969c.PNG"
		CaptionDelish = "*Бадриджани 175 г* \n" +
			"\n" +
			"Рулеты из жареных баклажанов с ореховой начинкой,\n" +
			" с добавлением граната и грецкого ореха \n" +
			"*Украшения:* веточка кинзы, грецкий орех, гранат.\n\n" +
			"*Состав:* Баклажаны, масло растительное,\n ореховый соус(" +
			"орех грецкий, семена подсолнечника, чеснок, лук репчатый,\n" +
			"масло подсолнечное, соль, уцхо-сунели, кориандр, перец красный,\n" +
			"винный уксус, сахар, паприка."

	case "adigay":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/784448c8ad7b47eefd214da9b0f7098a.PNG"
		CaptionDelish = "*Жареный адыгейский сыр с соусом наршараб 150 г*\n" +
			"Нежный адыгейский сыр, обжаренный в специях,\n" +
			"подаётся с гранатовым соусом наршараб.\n" +
			"*Украшения:* зерна граната и веточка кинзы.\n\n" +
			"*Состав:* Адыгейский сыр, масло подсолнечное, соль,\n" +
			" паприка, хмели-сунели, кориандр, перец черный,\n" +
			" уцхо-сенели, соус наршараб."
	case "vegetable":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/d34a9bdc475e09644e4da71be04bd99e.PNG"
		CaptionDelish = "*Овощной салат по-тбилисски 250 г*\n" +
			"Салат со свежими овощами, крымским луком, листом салата\n" +
			"и ореховой заправкой.\n" +
			"*Украшение:* лук ялтинский.\n\n" +
			"*Состав:* Помидор, огурец, перец болгарский, лист салат,\n" +
			"ореховая заправка (орех грецкий, семена подсолнечника,\n" +
			"масло подсолнечное, уксус винный, соль, сахар, уцхо-сунели,\n" +
			"кориандр, хмели-сунели, чеснок, аджика)"

	case "govyadina":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/9cdf8d3823e940970a732ff63a6755bc.PNG"
		CaptionDelish = "*Грузинский салат с говядиной 250 г*\n" +
			"Салат со свежими овощами, отварной говядиной, отварной\n" +
			"красной фасолью, гранатом, листом салата, ароматной зеленью\n" +
			"кинзы и ореховой заправкой.\n" +
			"*Украшения:* кинза, гранат.\n\n" +
			"*Состав:* Говядина, красная фасоль, помидор, перец болгарский,\n" +
			"лист салата, ореховая заправка ( орех грецкий, семена подсолничника,\n" +
			"аджика, уксус винный, масло подсолнечное, соль, сахар, чеснок, уцхо-сунели,\n" +
			"кориандр, хмели-сунели)."
	case "hrust":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/886aefda983a6451d388e907106b4dc4.PNG"
		CaptionDelish = "*Салат с хрустящими баклажанами и сыром 320 г*\n" +
			"Обжаренные баклажаны с хрустящей корочкой, адыгейский сыр, свежие помидоры и кинжа,\n" +
			"заправленные сладко-острым соусом и грецким орехом.\n\n" +
			"*Состав:* Баклажаны, масло подсолнечно, крахмал кукурузный, соль,\n" +
			"помидоры, зелень кинзы, соус чили сладкий, грецкий орех,\n" +
			"адыгейский сыр, сливки растительные, уцхо-сунели."
	case "warm":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/7e03c47e55f0fd65d02ab5bbe9f4be17.JPEG"
		CaptionDelish = "*Тёплый салат с говядиной 310 г*\n" +
			"Салат из обжаренного картофеля грибов, с отварной\n" +
			"говядиной, маринованными огурцами и квашеной капустой под\n" +
			"горчичной заправкой.\n\n" +
			"*Состав:* Картофель отварной в мундире, говядина отварная,\n" +
			"шампиньоны обжаренные, масло подсолнечное, огурцы маринованные\n" +
			"капуста квашеная, лук репчатый, уксус, соль, соус на салат Табуле\n" +
			"перец чили свежий."
	case "vinig":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/2069ba4e5731066d29f13a25d7830521.PNG"
		CaptionDelish = "*Винегрет с ореховой заправкой 250 г*\n" +
			"Салат из отварных овощей и капустой по-гурийски с ореховой заправкой.\n" +
			"*Украшение:* веточка кинзы\n\n" +
			"*Состав:* Свекла, морковь, картофель отварной, капуста маринованная\n" +
			"по-грузинский, горошек зеленый, заправка ореховая (орех грецкий, семена подсолнечника,\n" +
			"масло подсолнечное, уксус винный, соль, сахар, уцхо-сунели, кориандр, хмели-мунели,\n" +
			"чеснок, аджика), фасоль отварная, петрушка свежая, огурцы маринованные."
	case "potato":
		PhotoURL = "https://16a9564f-f8ec-42ba-a998-3027aa809e50.selstorage.ru/starikhinkalich-co/26130/14710/images/items/a2c5c36208f8c910b4d6a9b69b404fb0.PNG"
		CaptionDelish = "*Запечённый картофель с опятами и квашеной капустой.*\n" +
			"Салат из запеценного картофедя в специях, квашенной капусты с добавлением ароматного масла\n" +
			"и мариновынных опят.\n" +
			"*Украшения:* укроп, зелёный лук.\n\n" +
			"*Состав:* Капуста квашенная с луком(капуста квашеная, лук алтинский, масло растительное),\n" +
			"картофель в мундире запечённый(картофель отварной в мундире, масло подсолнечное,\n" +
			"соль, паприка, перец чёрный молотый, грибы опята, маринованные."

	}
	b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		Photo:     &models.InputFileString{Data: PhotoURL},
		Caption:   CaptionDelish,
		ParseMode: models.ParseModeMarkdownV1,
	})
}
