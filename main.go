package main

import (
	"context"
	"flag"
	"log"

	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"

	"github.com/sebasptsch/discraker/discord/commands"
	"github.com/sebasptsch/discraker/discord/utils"
	"github.com/sebasptsch/discraker/moonraker"
	"github.com/sourcegraph/jsonrpc2"
)

var commandDefinitions = []*discordgo.ApplicationCommand{
	&commands.EchoDefinition,
	&commands.WebcamsDefinition,
}

// Bot parameters
var (
	ConfigPath = flag.String("config", "~/printer_data/config/discraker.cfg", "Config file path")

	// AppId          = flag.String("app", "", "The application id")
)

func main() {
	flag.Parse()
	ReadConfig(*ConfigPath)
	discordSession, err := discordgo.New("Bot " + Config.Discord.Token)

	if err != nil {
		log.Panicf("Failed to create a discord session")
	}

	defer discordSession.Close()

	// 4. Set up an empty handler for incoming notifications/requests from server
	// The client needs a running background loop to continuously parse incoming reads
	handler := jsonrpc2.HandlerWithError(func(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (interface{}, error) {
		// Handle asynchronous server notifications or reverse calls here if required

		switch req.Method {
		// case "notify_active_spool_set":
		// 	{
		// 		channel, err := discordSession.UserChannelCreate(*ChannelId)
		// 		if err != nil {
		// 			log.Fatalln("Failed to create channel")
		// 		}

		// 		_, err = discordSession.ChannelMessageSend(channel.ID, "Spool updated")

		// 		if err != nil {
		// 			log.Fatalln("Failed to create message channel", err)
		// 		}
		// 		// s.ChannelMessageSend(*ChannelId, "Spool Updated")
		// 	}
		case "notify_proc_stat_update":
			{
				// This fires a lot even if I don't subscribe
			}
		default:
			{
				log.Printf("Received notification/request from server: %s", req.Method)
			}
		}
		return nil, nil
	})

	moonrakerSession, err := moonraker.New(Config.Moonraker.ConnectionURL, handler)

	if err != nil {
		log.Panicf("Failed to create moonraker session")
	}

	defer moonrakerSession.Close()

	// Interaction Handler
	discordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		data := i.ApplicationCommandData()

		switch data.Name {
		case "echo":
			commands.EchoHandler(s, i, utils.ParseOptions(data.Options))
		case "webcams":
			commands.WebcamsHandler(moonrakerSession, s, i)
		}

	})

	// Ready Handler
	discordSession.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s", r.User.String())
	})

	err = discordSession.Open() // Start the bot session
	if err != nil {
		log.Fatalf("could not open session: %s", err)
	}

	_, err = discordSession.ApplicationCommandBulkOverwrite(discordSession.State.Application.ID, Config.Discord.GuildID, commandDefinitions) // Send through the command definition
	if err != nil {
		log.Fatalf("could not register commands: %s", err)
	}

	_, err = moonrakerSession.PrinterInfo()

	if err != nil {
		log.Printf("Error querying printer info: %v", err)
	}

	// Handle Exit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
