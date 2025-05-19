package instructions

import (
	"github.com/df-mc/dragonfly/server/session"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type CameraInstruction interface {
	Packet(*session.Session) packet.Packet
}
