package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	h "github.com/holy-tech/discord-roulette/src/handlers"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid paramters: %v", err)
	}

	discord.AddHandler(h.Ready)
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = discord.Open()
	if err != nil {
		log.Fatalf("Could not open session: %v", err)
	}
	defer discord.Close()

	h.AppendHandler(discord, &h.RouletteHandle)
	h.AppendHandler(discord, &h.ShootHandle)
	h.AppendHandler(discord, &h.AcceptHandle)
	h.AppendHandler(discord, &h.DenyHandle)
	h.AppendHandler(discord, &h.GetGameHandle)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
