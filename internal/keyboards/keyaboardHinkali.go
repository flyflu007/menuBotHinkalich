package keyboards

import "github.com/go-telegram/bot/models"

func HinakliKeyboard() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{

				Text:         "Баранина и зелень",
				CallbackData: "hinkali_barzel",
			},
			{
				Text:         "Грибы",
				CallbackData: "hinkali_grib",
			},
			{

				Text:         "Свинина и говядина с зеленью",
				CallbackData: "hinkali_svingovzel",
			},
		},
		{
			{
				Text:         "Свинина и говяжина без зелени",
				CallbackData: "hinkali_svingov",
			},
			{
				Text:         "Три сыра",
				CallbackData: "hinkali_cheese",
			},
			{
				Text:         "Вишня и творог",
				CallbackData: "hinkali_vishtvor",
			},
		},
		{
			{
				Text:         "Говядина",
				CallbackData: "hinkali_gov",
			},
			{
				Text:         "Курица",
				CallbackData: "hinkali_kur",
			},
			{
				Text:         "Индейка",
				CallbackData: "hinkali_ind",
			},
		},
	}
}
