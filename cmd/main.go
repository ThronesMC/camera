package main

import (
	"github.com/ThronesMC/camera"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	chat.Global.Subscribe(chat.StdoutSubscriber{})

	c := server.UserConfig{}
	c.Network.Address = ":19132"
	c.Resources.Folder = "resources"
	c.Server.AuthEnabled = true

	conf, err := c.Config(slog.Default())
	if err != nil {
		panic(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	srv.Listen()
	for p := range srv.Accept() {
		camera.SendPresets(p)
		p.SetGameMode(world.GameModeCreative)
	}
}
