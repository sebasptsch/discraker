package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"

	"github.com/sebasptsch/discraker/discord/commands"
	"github.com/sebasptsch/discraker/discord/utils"
	"github.com/sebasptsch/discraker/moonraker"
	"github.com/sebasptsch/discraker/moonraker/structs"
	"github.com/sourcegraph/jsonrpc2"
)

var commandDefinitions = []*discordgo.ApplicationCommand{
	&commands.EchoDefinition,
	&commands.WebcamsDefinition,
	&commands.SnapshotDefinition,
}

// Bot parameters
var (
	ConfigPath = flag.String("config", "~/printer_data/config/discraker.cfg", "Config file path")

	// AppId          = flag.String("app", "", "The application id")
)

var version = "dev"

func main() {
	slog.Info(fmt.Sprintf("Running Discraker version: %s", version))
	flag.Parse()            // Parse CLI Flags
	ReadConfig(*ConfigPath) // Read the Config into memory

	discordSession, err := discordgo.New("Bot " + *Config.Discord.Token)
	if err != nil {
		panic(fmt.Errorf("failed to create discord session %w", err))
	}

	defer discordSession.Close()

	handler := jsonrpc2.HandlerWithError(func(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (interface{}, error) {
		// Handle asynchronous server notifications or reverse calls here if required

		switch req.Method {
		// case "notify_active_spool_set":
		// 	{
		// 		channel, err := discordSession.UserChannelCreate(*ChannelId)
		// 		if err != nil {
		// 			slog.Fatalln("Failed to create channel")
		// 		}

		// 		_, err = discordSession.ChannelMessageSend(channel.ID, "Spool updated")

		// 		if err != nil {
		// 			slog.Fatalln("Failed to create message channel", err)
		// 		}
		// 		// s.ChannelMessageSend(*ChannelId, "Spool Updated")
		// 	}
		case "notify_proc_stat_update":
			{
				// This fires a lot even if I don't subscribe
			}
		default:
			{
				slog.Warn(fmt.Sprintf("Received notification/request from server: %s", req.Method))
			}
		}
		return nil, nil
	})

	moonrakerConnectionParams := &moonraker.ConnectionParameters{
		HttpURL:   Config.Moonraker.HttpURL,
		SocketURL: Config.Moonraker.SocketURL,
		APIKey:    Config.Moonraker.APIKey,
	}

	moonrakerSession, err := moonraker.New(moonrakerConnectionParams, handler)

	if err != nil {
		panic(fmt.Errorf("failed to create moonraker session %w", err))
	}
	defer moonrakerSession.Close()

	// Interaction Handler
	discordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		data := i.ApplicationCommandData()
		var commandError error = nil
		switch data.Name {
		case "echo":
			commandError = commands.EchoHandler(s, i, utils.ParseOptions(data.Options))
		case "webcams":
			commandError = commands.WebcamsHandler(moonrakerSession, s, i)
		case "snapshot":
			commandError = commands.SnapshotHandler(moonrakerSession, s, i)
		}

		if commandError != nil {
			slog.Error(fmt.Sprintf("Command failed with error: %v", commandError))
		}

	})

	identifyReply, err := moonrakerSession.ServerConnectionIdentify(structs.ServerConnectionIdentifyParams{
		ClientName: "Discraker",
		Version:    version,
		Type:       "bot",
		URL:        "https://github.com/sebasptsch/discraker",
	})

	if err != nil {
		panic(fmt.Errorf("failed to get reply from moonraker identification request %w", err))
	}

	slog.Info(fmt.Sprintf("Connected to Moonraker with Connection ID: %d", identifyReply.ConnectionID))

	// Ready Handler
	discordSession.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		slog.Info(fmt.Sprintf("Logged in as %s", r.User.String()))
	})

	err = discordSession.Open() // Start the bot session
	if err != nil {
		panic(fmt.Errorf("failed to open discord session %w", err))
	}

	_, err = discordSession.ApplicationCommandBulkOverwrite(discordSession.State.Application.ID, *Config.Discord.GuildID, commandDefinitions) // Send through the command definition
	if err != nil {
		panic(fmt.Errorf("failed to publish/overwrite discord commands %w", err))
	}

	// Handle Exit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	slog.Info("Press Ctrl+C to exit")
	<-stop
}
