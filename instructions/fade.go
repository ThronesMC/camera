package instructions

import (
	"github.com/df-mc/dragonfly/server/session"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"image/color"
	"time"
)

type FadeCameraInstruction struct {
	// FadeInDuration is the time for the screen to fully fade in.
	FadeInDuration time.Duration
	// WaitDuration is time to wait before fading out.
	WaitDuration time.Duration
	// FadeOutDuration is the time for the screen to fully fade out.
	FadeOutDuration time.Duration

	// Colour is the colour of the screen to fade to. This only uses the red, green and blue components.
	Colour color.RGBA
}

func (i FadeCameraInstruction) Packet(*session.Session) packet.Packet {
	return &packet.CameraInstruction{
		Fade: protocol.Option(protocol.CameraInstructionFade{
			TimeData: protocol.Option(protocol.CameraFadeTimeData{
				FadeInDuration:  float32(i.FadeInDuration.Seconds()),
				WaitDuration:    float32(i.WaitDuration.Seconds()),
				FadeOutDuration: float32(i.FadeOutDuration.Seconds()),
			}),
			Colour: protocol.Option(i.Colour),
		}),
	}
}
