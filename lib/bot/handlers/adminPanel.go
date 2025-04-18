package handlers

import (
	"context"
	"strings"

	"github.com/JackBekket/hellper/lib/config"
	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

func (h *handlers) handleAdminStartАuthentication(ctx context.Context, tgb *bot.Bot, chatID int64) {
	log.Info().Int64("chat_id", chatID).Msg("Administrator authentication")

	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Enter your login..."}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	user.DialogStatus = statusAdminPanelEnterPassword
	h.cache.UpdateUser(user)
}

func (h *handlers) handleAdminPanelEnterPassword(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Enter your password..."}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	user.AdminLogin = update.Message.Text
	user.DialogStatus = statusAdminPanelCheckCredentials
	h.cache.UpdateUser(user)

}

func (h *handlers) handleAdminPanelCheckCredentials(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	l, p := user.AdminLogin, update.Message.Text
	if !checkAdminCredentialsn(l, p) {

	}

}

func checkAdminCredentialsn(login, password string) bool {
	l := strings.ReplaceAll(login, " ", "") == config.GetBot().Admin.Login
	p := strings.ReplaceAll(password, " ", "") == config.GetBot().Admin.Password

	return l && p
}
