package roulette

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	u "github.com/holy-tech/discord-roulette/src"
	"github.com/holy-tech/discord-roulette/src/data"
	db "github.com/holy-tech/discord-roulette/src/repo"
)

func GameStart(s *data.GameStatus) string {
	seatingName := []string{}
	for _, user := range s.Table.GetSeating() {
		seatingName = append(seatingName, user.Mention())
	}
	resp := fmt.Sprintf(
		`Preparing a %d-shooter with %d bullet(s). Prepare your self: %s`,
		s.Revolver.ChamberSize(), s.Revolver.GetNumBulletsLeft(), u.JoinStrings(", ", seatingName...),
	)
	if err := db.CreateGameDocument(s.Channel, s); err != nil {
		log.Printf("Error creating game document: %v", err)
		resp = fmt.Sprintf("Error: %v", err)
	}
	return resp
}

func ChallengeAccept(channel string, user *discordgo.User) string {
	var message string
	game, err := GetGame(channel)
	if err != nil {
		return user.Mention() + " Could not accept: " + err.Error()
	}
	game.Table.AcceptPlayer(user.ID)
	db.UpdateGameDocument(channel, game)
	if game.IsAccepted() {
		message += "\nIt is " + game.Table.GetCurrentPlayer().Mention() + " turn."
	}
	return user.Mention() + " has accepted!!\n" + message
}

func ChallengeDeny(channel string, user *discordgo.User) string {
	resp := "Putting gun away"
	if err := db.DeleteGameDocument(channel); err != nil {
		log.Printf("Error removing game: %v", err)
		resp = fmt.Sprintf("Error: %v", err)
	}
	return user.Mention() + " has denied!!\n" + resp
}
