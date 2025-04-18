package handlers

import "strings"

// Messages for the user
const (
	msgHello = "Hello! I am Hellper bot, choose the service you'd like to work with!"
	msgAwait = "Awaiting..."
	// Message with formatting for the user
	msgEnterAPIToken = "Your \"%s\" service, please enter your API token 🐱"
	msgChooseModel   = "Сhoose a model"
	//Message with formatting for the user
	msgSessionModelFormat = "Your session model: %s"
	//msgChooseLang         = "Choose a language or send 'Hello' in your desired language"
	msgFirstPrompt      = "Just say 'hello' in your language!"
	msgConnectingAINode = "Connecting to AI node..."
	msgHelpCommand      = "Authorize for additional commands: /help -- print this message, /restart -- restart session (if you want to switch between local-ai and openai chatGPT), /searchdoc -- searching documents, /rag -- process Retrival-Augmented Generation, /instruct -- use system promt template instead of langchain (higher priority, see examples), /image -- generate image ....all funcs are experimental so bot can halt and catch fire"
	msgAIclientFailure  = "An error has occured. In order to proceed we need to recreate client and initialize new session"
)

// const (
// 	langEnglish = "English"
// 	langRussian = "Russian"
// )

// Base prompts for the AI
const (
	basePromptLangRU         = "Привет, ты говоришь по-русски?"
	basePromptLangEN         = "Hi, Do you speak english?"
	basePromptRecognizeImage = "What's in the image?"
	basePromptGenerateImage  = "evangelion, neon, anime"
)

// Dialog status. The user's current position in the conversation with the bot
const (
	statusEnterYouAPIToken = iota + 1
	statusAuthMethodCallback
	statusLocalAIProviderCallback
	statusAPIToken
	//statusAIModelSelectionKeyboard
	statusAIModelSelectionChoiceCallback
	//statusConnectingToAIWithLangCallback
	statusConnectingToAIWithFirstPrompt
	statusStartDialogSequence

	// Dialog statuses in the admin panel
	statusAdminStartАuthentication
	statusAdminPanelEnterPassword
	statusAdminPanelCheckCredentials

	statusAdminPanelCallback
)

// Error messages for the user
const (
	errMsgFailedToGenerateImage = "Failed to generate image"
	errMsgFailedTrascribeVoice  = "Failed to transcribe the voice message"
	errMsgFailedRecognizeImage  = "Failed to recognize the image"
)

// Button text for keyboards
const (
	// func renderAdminPanelInlineKeyboard()
	kboardAdminPanelSettingsGroups = "Settings groups"
	kboardAdminPanelOther          = "Other"

	// func renderAIServicesInlineKeyboard()
	kboardAIServicesLocalAI = "LocalAI"
	kboardAIServicesOpenAI  = "OpenAI"

	// func renderAdminGroupInlineKeyboard()
	kboardAdminPanelNewToken    = "New token for this group"
	kboardAdminChangeServer     = "Change the Swarmind server"
	kboardAdminPanelBaseAIModel = "Change the bot's base AI model"
)

// Callback value constructor
func callbackText(buttomText string) string {
	noSpaces := strings.ReplaceAll(buttomText, " ", "")
	return "callback_" + noSpaces
}

// Default values for working with AI. Model names, URL endpointes
// const (
// 	aiStableDiffusionModel    = "stablediffusion"
// 	aiImageGenerationEndpoint = "/images/generations"

// 	aiBunnyLLAMAModel        = "bunny-llama-3-8b-v"
// 	aiImageRecognizeEndpoint = "/chat/completions"
// )

// Maximum number of results when searching for documents
const maxResultsForDoc = 3
