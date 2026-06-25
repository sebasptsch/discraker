package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sebasptsch/discraker/moonraker"
)

func WebcamsHandler(m *moonraker.Session, s *discordgo.Session, i *discordgo.InteractionCreate) {
	reply, err := m.WebcamsList()

	if err != nil {
		log.Panicln("Failed to get webcams")
	}

	var embeds []*discordgo.MessageEmbed

	for _, webcam := range reply.Webcams {
		var descriptionBuilder strings.Builder

		descriptionBuilder.WriteString(webcam.Service)

		webcamEmbed := discordgo.MessageEmbed{
			Title:       webcam.Name,
			URL:         webcam.StreamURL,
			Description: descriptionBuilder.String(),
			Fields:      []*discordgo.MessageEmbedField{{Name: "Service", Value: webcam.Service}, {Name: "Source", Value: webcam.Source}},
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
		log.Panicf("could not respond to interaction: %s", err)
	}

	fmt.Printf("Result successfully received: %+v\n", reply)
}

var WebcamsDefinition = discordgo.ApplicationCommand{
	Name:        "webcams",
	Description: "List webcams",
}
