package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sebasptsch/discraker/moonraker"
)

func WebcamsHandler(m *moonraker.Session, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	reply, err := m.ServerWebcamsList()

	if err != nil {
		return err
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
		return err
	}

	fmt.Printf("Result successfully received: %+v\n", reply)
	return nil
}

var WebcamsDefinition = discordgo.ApplicationCommand{
	Name:        "webcams",
	Description: "List webcams",
}
