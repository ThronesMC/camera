package main

import (
	"github.com/ThronesMC/camera"
	"github.com/ThronesMC/camera/instructions"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"time"
)

func init() {
	cmd.New("shake", "shakes it", nil, &ShakeCommand{})
}

type ShakeCommand struct {
}

func (s *ShakeCommand) Run(src cmd.Source, o *cmd.Output, _ *world.Tx) {
	if p, ok := src.(*player.Player); ok {
		camera.SendCameraInstruction(p, instructions.ShakeInstruction{
			Type:      instructions.CameraShakeTypeRotational,
			Intensity: 1.5,
			Duration:  time.Second * 2,
		})
		o.Print("successful!")
		return
	}
	o.Error("not a player")
}
