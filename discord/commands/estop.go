package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sebasptsch/discraker/moonraker"
)

func EStopHandler(m *moonraker.Session, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, err := m.PrinterEmergencyStop()

	if err != nil {
		return fmt.Errorf("unable to execute emergency stop %w", err)
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Printer Emergency Stopped",
			// Embeds:  embeds,
		},
	})

	if err != nil {
		return fmt.Errorf("failed to respond to interaction %w", err)
	}

	return nil
}

var EStopDefinition = discordgo.ApplicationCommand{
	Name:        "estop",
	Description: "Emergency stop",
}
