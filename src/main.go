package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dicord := connectToDiscord()
	registerCommands(dicord)
}

func registerCommands(discord *discordgo.Session) {
	_, err := discord.ApplicationCommandBulkOverwrite(1, 1, []*discordgo.ApplicationCommand{
		{
			Name:        "play",
			Description: "Plays a song",
		},
		{
			Name:        "clear",
			Description: "Clears the song queue",
		},
		{
			Name:        "skip",
			Description: "Skip the current song",
		},
		{
			Name:        "leave",
			Description: "Leave the current voice channel",
		},
	}) // TODO: fill these in
	if err != nil {
		log.Fatal("Error registering commands:", err)
	}
}

func connectToDiscord() *discordgo.Session {
	token := os.Getenv("TOKEN")
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
	}

	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildVoiceStates)

	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening connection to Discord", err)
	}

	return discord
}
