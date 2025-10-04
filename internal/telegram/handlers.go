package telegram

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/your-username/coffee-cups-system/internal/models"
)

// handleCoffee handles the /coffee command
func (b *Bot) handleCoffee(chatID int64, user *models.User, text string) {
	// Parse box ID from command: /coffee <box_id>
	parts := strings.Fields(text)
	if len(parts) != 2 {
		b.sendMessage(chatID, "Usage: /coffee <box_id>\nUse /boxes to see available boxes.")
		return
	}

	boxID, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		b.sendMessage(chatID, "Invalid box ID. Please provide a valid number.")
		return
	}

	// Log the coffee
	_, err = b.services.Coffee.LogCoffee(user.ID, uint(boxID))
	if err != nil {
		b.sendMessage(chatID, "Failed to log coffee: "+err.Error())
		return
	}

	// Get box info for response
	box, err := b.services.Box.GetBoxByID(uint(boxID))
	if err != nil {
		b.sendMessage(chatID, "Coffee logged successfully!")
		return
	}

	// Calculate remaining cups
	remaining, err := box.GetRemainingCups(b.services.Coffee.GetDB())
	if err != nil {
		b.sendMessage(chatID, "Coffee logged successfully!")
		return
	}

	msg := fmt.Sprintf("☕ Coffee logged successfully!\n\nBox: %s\nRemaining cups: %d", box.Name, remaining)
	b.sendMessage(chatID, msg)
}

// handleStatus handles the /status command
func (b *Bot) handleStatus(chatID int64, user *models.User) {
	// Get user's recent coffee logs
	logs, err := b.services.Coffee.GetUserCoffeeLogs(user.ID, 5)
	if err != nil {
		b.sendMessage(chatID, "Failed to get your coffee logs.")
		return
	}

	if len(logs) == 0 {
		b.sendMessage(chatID, "You haven't logged any coffee yet. Use /coffee <box_id> to log your first cup!")
		return
	}

	msg := "📊 Your recent coffee logs:\n\n"
	for _, log := range logs {
		msg += fmt.Sprintf("☕ %s - %s\n", log.Box.Name, log.LoggedAt.Format("2006-01-02 15:04"))
	}

	b.sendMessage(chatID, msg)
}

// handleBoxes handles the /boxes command
func (b *Bot) handleBoxes(chatID int64, _ *models.User) {
	boxes, err := b.services.Box.GetActiveBoxes()
	if err != nil {
		b.sendMessage(chatID, "Failed to get available boxes.")
		return
	}

	if len(boxes) == 0 {
		b.sendMessage(chatID, "No active boxes available.")
		return
	}

	msg := "📦 Available coffee boxes:\n\n"
	for _, box := range boxes {
		remaining, _ := box.GetRemainingCups(b.services.Coffee.GetDB())
		msg += fmt.Sprintf("ID: %d - %s\nPrice: $%.2f\nRemaining: %d/%d cups\n\n",
			box.ID, box.Name, box.Price, remaining, box.TotalCups)
	}

	msg += "Use /coffee <box_id> to log a coffee."
	b.sendMessage(chatID, msg)
}

// handleHelp handles the /help command
func (b *Bot) handleHelp(chatID int64) {
	msg := `🤖 Coffee Cups System Bot

Available commands:
/start - Start using the bot
/coffee <box_id> - Log a coffee consumption
/status - View your recent coffee logs
/boxes - View available coffee boxes
/help - Show this help message

How it works:
1. Use /boxes to see available coffee boxes
2. Use /coffee <box_id> to log when you take a coffee
3. The system automatically calculates your share of the cost
4. Use /status to see your consumption history

Happy coffee drinking! ☕`

	b.sendMessage(chatID, msg)
}

// sendMessage sends a message to a chat
func (b *Bot) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"

	if _, err := b.api.Send(msg); err != nil {
		// Log error but don't crash
		fmt.Printf("Failed to send message: %v\n", err)
	}
}
