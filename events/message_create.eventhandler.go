package events

import (
	"github.com/bwmarrin/discordgo"
	"lets-try-go/commands"
	"lets-try-go/utils"
	"strings"
)


var (
	commandList = []commands.Command{
		commands.Ping(),
		commands.Help(),
	}
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

			utils.Debug("Trying to get command: " + commandName)

			ExecuteCommand(commandName, session, event, commandArgs)
		}
	}}
	return
}

// everything here because import cycle is not allowed here :(

// GetCommand search for a matching command
// from the command name in the parameter
func GetCommand(commandName string) (command commands.Command) {
	utils.Debug("Trying to get command: " + commandName)
	for _, element := range commandList {
		if element.Name == strings.ToLower(commandName) {
			command = element
		}
	}
	return
}

// ExecuteCommand gets the command from the parameter,
// then execute it if exists
func ExecuteCommand(commandName string, session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
	cmd := GetCommand(commandName)
	if cmd.Exec == nil {
		return
	}
	cmd.Exec(session, event, args)
}

