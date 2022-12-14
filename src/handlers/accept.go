package handlers

import (
	"github.com/bwmarrin/discordgo"
	r "github.com/holy-tech/discord-roulette/src/roulette"
)

var AcceptHandle = Handler{
	CommandSpecs: &discordgo.ApplicationCommand{
		Name:                     "roulette-accept",
		Description:              "Accept roulette match",
		DefaultMemberPermissions: &defaultAdmin,
	},
	CommandHandler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		user := i.Member.User
		channel := i.ChannelID
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: r.ChallengeAccept(channel, user),
			},
		})
	},
}

var DenyHandle = Handler{
	CommandSpecs: &discordgo.ApplicationCommand{
		Name:                     "roulette-deny",
		Description:              "Deny roulette match",
		DefaultMemberPermissions: &defaultAdmin,
	},
	CommandHandler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		challenger := i.Member.User
		channel := i.ChannelID
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: r.ChallengeDeny(channel, challenger),
			},
		})
	},
}
