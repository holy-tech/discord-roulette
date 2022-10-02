package data

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

var (
	DefaultOpponents             map[string]Player = map[string]Player{}
	DefaultGameAccepted          bool              = false
	DefaultSpinChamberRule       bool              = false
	DefaultSpinChamberOnShotRule bool              = false
	DefaultReplaceBulletRule     bool              = false
	DefaultChannel               string            = ""
)

type Player struct {
	discordgo.User
	Accepted string `json:"accepted"`
}

type GameStatus struct {
	Opponents             map[string]Player
	TableState            TableState
	Revolver              Revolver
	GameAccepted          bool   `json:"game_accepted,omitempty"`
	SpinChamberRule       bool   `json:"spin_chamber,omitempty"`
	SpinChamberOnShotRule bool   `json:"spin_chamber_on_shot,omitempty"`
	ReplaceBulletRule     bool   `json:"replace_bullet,omitempty"`
	Channel               string `json:"channel,omitempty"`
}

var DefaultGameStatus GameStatus = GameStatus{
	DefaultOpponents,
	DefaultTableState,
	DefaultRevolver,
	DefaultGameAccepted,
	DefaultSpinChamberRule,
	DefaultSpinChamberOnShotRule,
	DefaultReplaceBulletRule,
	DefaultChannel,
}

func (s *GameStatus) Shoot(user *discordgo.User) (bool, error) {
	currPlayer := s.TableState.GetCurrentPlayer()
	if user.ID != currPlayer {
		return false, errors.New("it is not your turn")
	}
	shot := s.Revolver.Chambers[s.Revolver.CurrentChamber]
	s.Revolver.SetNextChamber()
	s.TableState.SetNextPlayer()
	if shot {
		s.TableState.Losers = append(s.TableState.Losers, user.ID)
		delete(s.Opponents, user.ID)
		s.TableState.RemovePlayer(user.ID)
	}
	s.Revolver.ClearChamber(shot)
	if s.Revolver.NumBulletsLeft <= 0 {
		s.Revolver.SpinChamber()
	}
	return shot, nil
}