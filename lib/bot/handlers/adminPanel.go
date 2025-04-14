package handlers

import (
	"context"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// A wrapper around the "admin_panel" command, needed for re-invocation in case the admin password was entered incorrectly.
// This wrapper is used in `routerMessage` to allow calling the function based on the state status.
func (h *handlers) handleAdminАuthentication(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	h.cmdAdminPanel(ctx, tgb, update.Message.From.ID)
}

func adminАuthentication(user *database.User, tgb *bot.Bot) bool
