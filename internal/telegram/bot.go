package telegram

import (
	"context"
	"fmt"
	"strings"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/your-username/coffee-cups-system/internal/config"
	"github.com/your-username/coffee-cups-system/internal/models"
	"github.com/your-username/coffee-cups-system/internal/services"
)

// Bot represents the Telegram bot
type Bot struct {
	api      *telegram.BotAPI
	services *services.Services
	config   config.TelegramConfig
	logger   interface{}
}

// New creates a new Telegram bot instance
func New(cfg config.TelegramConfig, services *services.Services, logger interface{}) (*Bot, error) {
	bot, err := telegram.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	bot.Debug = cfg.Debug

	return &Bot{
		api:      bot,
		services: services,
		config:   cfg,
		logger:   logger,
	}, nil
}

// Start starts the bot and begins listening for updates
func (b *Bot) Start(ctx context.Context) error {
	u := telegram.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case update := <-updates:
			if update.Message != nil {
				b.handleMessage(update.Message)
			}
		}
	}
}

// handleMessage handles incoming messages
func (b *Bot) handleMessage(message *telegram.Message) {
	chatID := message.Chat.ID
	text := message.Text

	// Create or update user
	user, err := b.services.User.CreateOrUpdateUser(
		int64(message.From.ID),
		message.From.UserName,
		message.From.FirstName,
		message.From.LastName,
	)
	if err != nil {
		b.sendMessage(chatID, "Sorry, there was an error processing your request.")
		return
	}

	// Handle commands
	switch {
	case strings.HasPrefix(text, "/start"):
		b.handleStart(chatID, user)
	case strings.HasPrefix(text, "/coffee"):
		b.handleCoffee(chatID, user, text)
	case strings.HasPrefix(text, "/status"):
		b.handleStatus(chatID, user)
	case strings.HasPrefix(text, "/boxes"):
		b.handleBoxes(chatID, user)
	case strings.HasPrefix(text, "/help"):
		b.handleHelp(chatID)
	default:
		b.sendMessage(chatID, "I don't understand that command. Use /help to see available commands.")
	}
}

// handleStart handles the /start command
func (b *Bot) handleStart(chatID int64, user *models.User) {
	msg := fmt.Sprintf("Welcome %s! I'm your coffee tracking bot. Use /help to see available commands.", user.FirstName)
	b.sendMessage(chatID, msg)
}

// ... (other handler methods would be implemented here)
