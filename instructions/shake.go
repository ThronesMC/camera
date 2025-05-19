package instructions

import (
	"github.com/df-mc/dragonfly/server/session"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"time"
)

const (
	CameraShakeTypePositional uint8 = iota
	CameraShakeTypeRotational
)

type ShakeInstruction struct {
	// Type is the type of shake, and is one of the constants listed above. The different type affects how
	// the shake looks in game.
	Type uint8
	// Intensity is the intensity of the shaking. The client limits this value to 4, so anything higher may
	// not work.
	Intensity float32
	// Duration is for how long the camera will shake for.
	Duration time.Duration
}

func (i ShakeInstruction) Packet(*session.Session) packet.Packet {
	if i.Type < CameraShakeTypePositional || i.Type > CameraShakeTypeRotational {
		panic("type is out of range")
	}
	return &packet.CameraShake{
		Intensity: i.Intensity,
		Duration:  float32(i.Duration.Seconds()),
		Type:      i.Type,
		Action:    packet.CameraShakeActionAdd,
	}
}

type StopShakeInstruction struct {
}

func (i StopShakeInstruction) Packet(*session.Session) packet.Packet {
	return &packet.CameraShake{
		Action: packet.CameraShakeActionStop,
	}
}
