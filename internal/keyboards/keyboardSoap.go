package keyboards

import "github.com/go-telegram/bot/models"

func SoupKeyboard() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{

				Text:         "Харчо с говядиной",
				CallbackData: "soup_harcho",
			},
			{
				Text:         "Крем-суп чкмерули",
				CallbackData: "soup_cremsoup",
			},
		},
		{
			{

				Text:         "Чихиртма с курицей",
				CallbackData: "soup_chihir",
			},
			{
				Text:         "Грибная юшка",
				CallbackData: "soup_grib",
			},
		},
	}
}
