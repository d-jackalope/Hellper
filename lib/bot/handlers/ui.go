package handlers

import (
	"strconv"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot/models"
)

// Render AI models menu with Inline Keyboard
func renderAIModelsInlineKeyboard(aiModelsList []string) models.InlineKeyboardMarkup {
	buttons := [][]models.InlineKeyboardButton{}
	for _, model := range aiModelsList {
		buttons = append(buttons, []models.InlineKeyboardButton{
			{
				Text:         model,
				CallbackData: "model_" + model,
			},
		})
	}

	return models.InlineKeyboardMarkup{
		InlineKeyboard: buttons,
	}
}

// Render LocalAI Providers menu with Inline Keyboard
func renderLocalAIProvidersInlineKeyboard(aiProviderList []string) models.InlineKeyboardMarkup {
	buttons := [][]models.InlineKeyboardButton{}
	for _, provider := range aiProviderList {
		buttons = append(buttons, []models.InlineKeyboardButton{
			{
				Text:         provider,
				CallbackData: callbackText(provider),
			},
		})
	}

	return models.InlineKeyboardMarkup{
		InlineKeyboard: buttons,
	}
}

func renderAIServicesInlineKeyboard() models.InlineKeyboardMarkup {
	return models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: kboardAIServicesLocalAI, CallbackData: strconv.Itoa(database.AuthMethodLocalAI)},
				{Text: kboardAIServicesOpenAI, CallbackData: strconv.Itoa(database.AuthMethodOpenAI)},
			},
		},
	}
}

func renderAdminPanelInlineKeyboard() models.InlineKeyboardMarkup {
	return models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: kboardAdminPanelSettingsGroups, CallbackData: callbackText(kboardAdminPanelSettingsGroups)},
				{Text: kboardAdminPanelOther, CallbackData: callbackText(kboardAdminPanelOther)},
			},
		},
	}
}

func renderAdminSettingsGroupInlineKeyboard(groupList []string) models.InlineKeyboardMarkup {
	buttons := [][]models.InlineKeyboardButton{}
	for _, group := range groupList {
		buttons = append(buttons, []models.InlineKeyboardButton{
			{
				Text:         group,
				CallbackData: callbackText(group),
			},
		})
	}

	return models.InlineKeyboardMarkup{
		InlineKeyboard: buttons,
	}
}

func renderAdminGroupInlineKeyboard() models.InlineKeyboardMarkup {
	return models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: kboardAdminPanelNewToken, CallbackData: callbackText(kboardAdminPanelNewToken)},
				{Text: kboardAdminChangeServer, CallbackData: callbackText(kboardAdminChangeServer)},
				{Text: kboardAdminPanelBaseAIModel, CallbackData: callbackText(kboardAdminPanelBaseAIModel)},
			},
		},
	}
}

// Render Language menu with Inline Keyboard
// func renderLangInlineKeyboard() models.InlineKeyboardMarkup {
// 	return models.InlineKeyboardMarkup{
// 		InlineKeyboard: [][]models.InlineKeyboardButton{
// 			{
// 				{Text: langEnglish, CallbackData: langEnglish},
// 				{Text: langRussian, CallbackData: langRussian},
// 			},
// 		},
// 	}
// }
