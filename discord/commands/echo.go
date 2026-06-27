package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sebasptsch/discraker/discord/utils"
)

func EchoHandler(s *discordgo.Session, i *discordgo.InteractionCreate, opts utils.OptionMap) error {
	builder := new(strings.Builder)
	if v, ok := opts["author"]; ok && v.BoolValue() {
		author := utils.InteractionAuthor(i.Interaction)
		builder.WriteString("**" + author.String() + "** says: ")
	}
	builder.WriteString(opts["message"].StringValue())

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: builder.String(),
		},
	})

	return err
}

var EchoDefinition = discordgo.ApplicationCommand{
	Name:        "echo",
	Description: "Say something through a bot",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "message",
			Description: "Contents of the message",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
		{
			Name:        "author",
			Description: "Whether to prepend message's author",
			Type:        discordgo.ApplicationCommandOptionBoolean,
		},
	},
}
