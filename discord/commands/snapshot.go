package commands

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sebasptsch/discraker/moonraker"
)

func SnapshotHandler(m *moonraker.Session, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	reply, err := m.WebcamsList()

	if err != nil {
		return err
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Taking Snapshot",
		},
	})

	if err != nil {
		return err
	}

	var snapshots []*discordgo.File

	var builder strings.Builder
	for _, webcam := range reply.Webcams {
		builder.WriteString("** ")
		builder.WriteString(":camera_with_flash:")
		builder.WriteString(webcam.Name)
		builder.WriteString(" **")
		builder.WriteString("\n")
		if len(webcam.SnapshotURL) > 0 {
			resp, err := http.Get(webcam.SnapshotURL)
			if err != nil {
				break
			}
			defer resp.Body.Close()

			imageUrl, err := url.Parse(webcam.SnapshotURL)

			if err != nil {
				break
			}

			snapshots = append(snapshots, &discordgo.File{
				Name:        imageUrl.Path,
				ContentType: resp.Header.Get("Content-Type"),
				Reader:      resp.Body,
			})
		}

	}

	content := builder.String()

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &content,
		Files:   snapshots,
	})

	if err != nil {
		slog.Error(fmt.Sprintf("could not respond to interaction: %s", err))
		return err
	}

	slog.Info(fmt.Sprintf("Result successfully received: %+v\n", reply))
	return nil
}

var SnapshotDefinition = discordgo.ApplicationCommand{
	Name:        "snapshot",
	Description: "Snapshot",
}
