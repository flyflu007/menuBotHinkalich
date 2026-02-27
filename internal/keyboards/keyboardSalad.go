package keyboards

import "github.com/go-telegram/bot/models"

func SaladKeyboard() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{

				Text:         "Бандриджани",
				CallbackData: "salad_badrid",
			},
			{
				Text:         "Жареный адыгейский сыр с соусом наршараб",
				CallbackData: "salad_adigay",
			},
			{

				Text:         "Овощной салат по-тбилисски",
				CallbackData: "salad_vegetable",
			},
		},
		{
			{
				Text:         "Грузинский салад с говядиной",
				CallbackData: "salad_govyadina",
			},
			{
				Text:         "Салат с хрустящими баклажанами и сыром",
				CallbackData: "salad_hrust",
			},
			{
				Text:         "Тёплый салат с говядиной",
				CallbackData: "salad_warm",
			},
		},
		{
			{
				Text:         "Винегрет с ореховой заправкой",
				CallbackData: "salad_vinig",
			},
			{
				Text:         "Запеченный картофель с опятами и квашеной капустой",
				CallbackData: "salad_potato",
			},
		},
	}
}
