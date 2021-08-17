package commands

import (
"fmt"
"github.com/bwmarrin/discordgo"
"lets-try-go/utils"
"time"
)

func Ping() (command Command) {
	command = Command{Name: "ping", Description: "Calculate the bot's current ping",
		Exec: func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
			utils.Debug("Received command: Ping")

			beforeTime := time.Now()
			// send an embed
			msg, _ := session.ChannelMessageSendEmbed(event.ChannelID, utils.NewEmbed().SetDescription("Calculating my ping...").SetColor(3215252).MessageEmbed)
			afterTime := time.Now().Sub(beforeTime).String()
			time.Sleep(2 * time.Second)

			// edit the embed
			_ , err := session.ChannelMessageEditEmbed(event.ChannelID, msg.ID, utils.NewEmbed().SetDescription("My Ping: " + afterTime).SetColor(1283252).MessageEmbed)

			// check for errors
			if err != nil {
				utils.Error("An error occurred while sending an embed!")
				utils.Error(fmt.Sprint(err))
			}
		},
	}
	return
}
