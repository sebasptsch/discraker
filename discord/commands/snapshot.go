package commands

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
	moonrakerclient "github.com/sebasptsch/discraker/moonraker-client"
)

func SnapshotHandler(m *moonrakerclient.Session, s *discordgo.Session, i *discordgo.InteractionCreate) error {

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	if err != nil {
		return fmt.Errorf("failed to reply to interaction %w", err)
	}

	reply, err := m.ServerWebcamsList()

	if err != nil {
		return fmt.Errorf("unable to fetch webcam list %w", err)
	}

	var files = []*discordgo.File{}
	var embeds = []*discordgo.MessageEmbed{}

	for _, webcam := range reply.Webcams {
		embed := discordgo.MessageEmbed{}

		if webcam.SnapshotURL != nil {
			resp, err := http.Get(*webcam.SnapshotURL)
			if err != nil {
				return fmt.Errorf("failed to get screenshot from webcam %s: %w", webcam.Name, err)
			}
			defer resp.Body.Close()

			imageUrl, err := url.Parse(*webcam.SnapshotURL)

			if err != nil {
				return fmt.Errorf("failed to parse webcam snapshot url %s: %w", webcam.Name, err)
			}

			ext := filepath.Ext(imageUrl.Path)

			filename := fmt.Sprintf("%s%s", webcam.Name, ext)

			files = append(files, &discordgo.File{
				Name:        filename,
				ContentType: resp.Header.Get("Content-Type"),
				Reader:      resp.Body,
			})

			slog.Debug(filename)

			embed.Image = &discordgo.MessageEmbedImage{
				URL: fmt.Sprintf("attachment://%s", filename),
			}
		}

		embed.Title = webcam.Name

		embeds = append(embeds, &embed)
	}

	msg := "\u200b"

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &msg,
		Files:   files,
		Embeds:  &embeds,
	})

	if err != nil {
		return fmt.Errorf("failed to edit interaction response %w", err)
	}

	return nil
}

var SnapshotDefinition = discordgo.ApplicationCommand{
	Name:        "snapshot",
	Description: "Snapshot",
}
