package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	discord := connectToDiscord()
	initCommands(discord)

	discord.AddHandler(func(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
		data := interaction.ApplicationCommandData()
		switch data.Name {
		case "play":
			var response string

			status, title := play("TODO: song name")
			if status {
				response = "Song added to queue: " + title
			} else {
				response = "Song not added to queue"
			}

			err := discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: response,
				},
			})
			if err != nil {
				log.Fatal("Error adding command handler:", err)
			}
		}
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	err = discord.Close()
	if err != nil {
		log.Fatal("Error closing Discord:", err)
	}
}

func initCommands(discord *discordgo.Session) {
	_, err := discord.ApplicationCommandBulkOverwrite(discord.State.Application.ID, "", []*discordgo.ApplicationCommand{
		{
			Name:        "play",
			Description: "Plays a song",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "song",
					Description: "Song title or URL",
					Required:    true,
				},
			},
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
	})
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

	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)

	err = discord.Open()
	if err != nil {
		log.Fatal("Error connecting to Discord:", err)
	}

	return discord
}
