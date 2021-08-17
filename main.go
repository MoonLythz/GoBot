package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"lets-try-go/managers"
	"lets-try-go/utils"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	token = "TOKEN"
)

func main() {

	utils.Log("Starting the bot")
	beforeTime := time.Now()
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		utils.Error("An error occurred while starting the bot!")
		utils.Error(fmt.Sprint(err))
		return
	}
	
	managers.EventManager(discord)

	utils.Log("Opening a websocket connection to discord!")
	err = discord.Open()
	if err != nil {
		utils.Error("An error occurred while opening a websocket connection!")
		utils.Error(fmt.Sprint(err))
		return
	}
	afterTime := time.Now().Sub(beforeTime)
	utils.Log("Done! Bot is now running!")
	utils.Debug("Took " + strconv.FormatInt(afterTime.Milliseconds(), 10) + "ms")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err = discord.Close()
	utils.Log("Closing discord client")
	if err != nil {
		utils.Error("An error occurred while closing the discord client!")
		utils.Error(fmt.Sprint(err))
	}
}