package managers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"lets-try-go/events"
	"lets-try-go/utils"
)

var (
	eventHandlers = []events.EventHandler{
		events.MessageCreate(),
	}
)

func EventManager(session *discordgo.Session) {
	// register all the events
	utils.Log("Registering event listeners!")
	for _, element := range eventHandlers {
		utils.Debug("Registered " + fmt.Sprint(element))
		session.AddHandler(element.Handler)
	}
}
