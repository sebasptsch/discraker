package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sebasptsch/discraker/moonraker"
)

func StatusHandler(m *moonraker.Session, s *discordgo.Session, i *discordgo.InteractionCreate) error {
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

var StatusDefinition = discordgo.ApplicationCommand{
	Name:        "status",
	Description: "Get the current status of the printer.",
}
