package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	moonrakerclient "github.com/sebasptsch/discraker/moonraker-client"
)

func WebcamsHandler(m *moonrakerclient.Session, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	reply, err := m.ServerWebcamsList()

	if err != nil {
		return fmt.Errorf("unable to fetch webcam list %w", err)
	}

	var embeds []*discordgo.MessageEmbed

	for _, webcam := range reply.Webcams {
		var descriptionBuilder strings.Builder

		descriptionBuilder.WriteString(webcam.Service)

		var fields = []*discordgo.MessageEmbedField{
			{Name: "Service", Value: webcam.Service, Inline: true},
			{Name: "Source", Value: webcam.Source, Inline: true},
			{Name: "Enabled", Value: fmt.Sprintf("%t", webcam.Enabled), Inline: true},
		}

		webcamEmbed := discordgo.MessageEmbed{
			Title:  webcam.Name,
			Fields: fields,
		}

		embeds = append(embeds, &webcamEmbed)
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "# Webcams",
			Embeds:  embeds,
		},
	})

	if err != nil {
		return fmt.Errorf("failed to respond to interaction %w", err)
	}

	return nil
}

var WebcamsDefinition = discordgo.ApplicationCommand{
	Name:        "webcams",
	Description: "List webcams",
}
