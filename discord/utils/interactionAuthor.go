package utils

import "github.com/bwmarrin/discordgo"

func InteractionAuthor(i *discordgo.Interaction) *discordgo.User {
	if i.Member != nil {
		return i.Member.User
	}
	return i.User
}
