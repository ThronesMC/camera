package instructions

import (
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"

	_ "unsafe"
)

type TargetCameraInstruction struct {
	// CenterOffset is the offset from the center of the entity that the camera should target.
	CenterOffset *mgl32.Vec3
	// EntityHandle is the entity handle that the camera should target.
	EntityHandle *world.EntityHandle
}

func (i TargetCameraInstruction) Packet(s *session.Session) packet.Packet {
	var centerOffset protocol.Optional[mgl32.Vec3]
	if i.CenterOffset != nil {
		centerOffset = protocol.Option(*i.CenterOffset)
	}

	return &packet.CameraInstruction{
		Target: protocol.Option(protocol.CameraInstructionTarget{
			CenterOffset:   centerOffset,
			EntityUniqueID: int64(session_handleRuntimeID(s, i.EntityHandle)),
		}),
	}
}

//go:linkname session_handleRuntimeID github.com/df-mc/dragonfly/server/session.(*Session).handleRuntimeID
func session_handleRuntimeID(*session.Session, *world.EntityHandle) uint64
