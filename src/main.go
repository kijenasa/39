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
}
