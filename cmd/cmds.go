package main

import (
	"github.com/ThronesMC/camera"
	"github.com/ThronesMC/camera/instructions"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"time"
)

func init() {
	cmd.Register(cmd.New("shake", "shakes it", nil, NewCameraCommand(instructions.ShakeInstruction{
		Type:      instructions.CameraShakeTypeRotational,
		Intensity: 1.5,
		Duration:  time.Second * 2,
	})))
	cmd.Register(cmd.New("free", "birb moment", nil, NewCameraCommand(instructions.SetCameraInstruction{
		Preset: instructions.FreeCameraPreset,
		Ease: protocol.Option(instructions.CameraEase{
			Type:     instructions.EasingTypeInOutBounce,
			Duration: time.Second,
		}),
		Position: protocol.Option(mgl32.Vec3{0, 30, 0}),
		Facing:   protocol.Option(mgl32.Vec3{0, 0, 0}),
	})))
	cmd.Register(cmd.New("fades", "raahhh im cool", nil, NewCameraCommand(instructions.FadeCameraInstruction{
		FadeInDuration:  500 * time.Millisecond,
		WaitDuration:    1 * time.Second,
		FadeOutDuration: 250 * time.Millisecond,
		Colour:          item.ColourYellow().RGBA(),
	})))
	cmd.Register(cmd.New("clear", "back to boring", nil, NewCameraCommand(instructions.ClearCameraInstruction{})))
}

func NewCameraCommand(instruction instructions.CameraInstruction) *CameraCommand {
	return &CameraCommand{instruction: instruction}
}

type CameraCommand struct {
	instruction instructions.CameraInstruction
}

func (c CameraCommand) Run(src cmd.Source, o *cmd.Output, _ *world.Tx) {
	if p, ok := src.(*player.Player); ok {
		camera.SendCameraInstruction(p, c.instruction)
		o.Print("successful!")
		return
	}
	o.Error("not a player")
}
