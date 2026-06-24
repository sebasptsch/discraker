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

	var builder strings.Builder
	for _, webcam := range reply.Webcams {
		builder.WriteString(webcam.Name)
		builder.WriteString("\n")
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: builder.String(),
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
