package events

import (
	"github.com/bwmarrin/discordgo"
)

// an "abstract class" for event handlers
type EventHandler struct {
	Handler func(session *discordgo.Session, event interface{})
}
