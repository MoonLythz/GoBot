package events

import (
	"github.com/bwmarrin/discordgo"
	"lets-try-go/commands"
	"lets-try-go/utils"
	"strings"
)

// MessageCreate listens for the MessageCreate event
func MessageCreate() (eventHandler EventHandler) {
	eventHandler = EventHandler{Handler: func(session *discordgo.Session, event interface{}) {
		// check if the event is a MessageCreate event
		if event, ok := event.(*discordgo.MessageCreate); ok {
			utils.Debug("Received event: MessageCreate")
			// verification stuff
			if event.ID == session.State.User.ID { return }
			if !strings.HasPrefix(event.Content, "mb!") { return }

			// split the message into command name, and args
			commandFull := strings.Split(event.Content, "mb!")[1]
			if commandFull == "" { return }

			commandName := strings.Split(commandFull, " ")[0]

			commandArgs := strings.Split(commandFull, " ")
			// if the message has args
			if len(commandArgs) > 1 {
				emptyArray := make([]string, 0)
				emptyArray = append(emptyArray, commandArgs[:1]...)
				commandArgs = append(emptyArray, commandArgs[1+1:]...)
			}
7
			utils.Debug("Trying to get command: " + commandName)

			commands.ExecuteCommand(commandName, session, event, commandArgs)
		}
	}}
	return
}

// everything here because import cycle is not allowed here :(
