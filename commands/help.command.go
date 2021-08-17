package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"lets-try-go/utils"
)

func Help() (command Command) {
	command = Command{Name: "help", Description: "List of the bot's available commands",
		Exec: func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
			utils.Debug("Received command: Help")

			// send an embed
			_, err :=session.ChannelMessageSendEmbed(event.ChannelID,
				utils.NewEmbed().SetDescription(
					"`ping` - Calculate the current bot's ping\n" +
					" `help` - List of the bot's available commands").
				SetTitle("Command List").SetColor(1283252).MessageEmbed)

			// check for errors
			if err != nil {
				utils.Error("An error occurred while sending an embed!")
				utils.Error(fmt.Sprint(err))
			}
		},
	}
	return
}