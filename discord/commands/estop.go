package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sebasptsch/discraker/moonraker"
)

func EStopHandler(m *moonraker.Session, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, err := m.PrinterEmergencyStop()

	if err != nil {
		return err
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Printer Emergency Stopped",
			// Embeds:  embeds,
		},
	})

	return err
}

var EStopDefinition = discordgo.ApplicationCommand{
	Name:        "estop",
	Description: "Emergency stop",
}
