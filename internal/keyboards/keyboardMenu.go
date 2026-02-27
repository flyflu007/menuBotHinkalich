package keyboards

import "github.com/go-telegram/bot/models"

func MenuKeyboard() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{

				Text:         "Салаты",
				CallbackData: "salad",
			},
			{
				Text:         "Супы",
				CallbackData: "soup",
			},
			{

				Text:         "Основные блюда",
				CallbackData: "main",
			},
		},
		{
			{
				Text:         "Хинкали",
				CallbackData: "hinkali",
			},
			{
				Text:         "Соусы",
				CallbackData: "sauce",
			},
			{
				Text:         "Мучное",
				CallbackData: "muchnoe",
			},
		},
		{
			{
				Text:         "Сладкое",
				CallbackData: "sweet",
			},
		},
	}
}
