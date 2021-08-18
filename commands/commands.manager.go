package commands

import (
	"github.com/bwmarrin/discordgo"
	"lets-try-go/utils"
	"strings"
)

var (
	commandList = []Command{
		Ping(),
		Help(),
	}
)

// GetCommand search for a matching command
// from the command name in the parameter
func GetCommand(commandName string) (command Command) {
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