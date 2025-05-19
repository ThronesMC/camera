package instructions

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type ClearCameraInstruction struct {
}

func (i ClearCameraInstruction) Packet() packet.Packet {
	return &packet.CameraInstruction{
		Set:    protocol.Optional[protocol.CameraInstructionSet]{},
		Clear:  protocol.Option(true),
		Fade:   protocol.Optional[protocol.CameraInstructionFade]{},
		Target: protocol.Optional[protocol.CameraInstructionTarget]{},
	}
}
