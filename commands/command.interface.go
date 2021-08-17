package commands

import "github.com/bwmarrin/discordgo"

// an "abstract class" for commands
type Command struct {
	Name        string
	Description string
	Exec        func(session *discordgo.Session, event *discordgo.MessageCreate, args []string)
}
