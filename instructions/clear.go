package instructions

import (
	"github.com/df-mc/dragonfly/server/session"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type ClearCameraInstruction struct {
}

func (i ClearCameraInstruction) Packet(*session.Session) packet.Packet {
	return &packet.CameraInstruction{
		Clear: protocol.Option(true),
	}
}
