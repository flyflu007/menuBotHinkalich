package keyboards

import "github.com/go-telegram/bot/models"

func MainKeyboard() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{

				Text:         "Плов с говядиной",
				CallbackData: "main_plov",
			},
			{
				Text:         "Чахохбили из курицы",
				CallbackData: "main_chahoh",
			},
			{

				Text:         "Мини-хинкали с курицей в сырном соусе",
				CallbackData: "main_minihin",
			},
		},
		{
			{
				Text:         "Оджахури со свининой",
				CallbackData: "main_odjahur",
			},
			{
				Text:         "Лагман с говядиной",
				CallbackData: "main_lagman",
			},
			{
				Text:         "Долма",
				CallbackData: "main_dolma",
			},
		},
		{
			{
				Text:         "Стейк из капусты с ореховым соусом",
				CallbackData: "main_stake",
			},
		},
	}
}
