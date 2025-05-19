package camera

import (
	"github.com/ThronesMC/camera/instructions"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	_ "unsafe"
)

func SendPresets(p *player.Player) {
	s := player_session(p)
	session_writePacket(s, &packet.CameraPresets{
		Presets: instructions.Presets,
	})
}

func SendCameraInstruction(p *player.Player, i instructions.CameraInstruction) {
	s := player_session(p)
	session_writePacket(s, i.Packet(s))
}

//go:linkname player_session github.com/df-mc/dragonfly/server/player.(*Player).session
func player_session(*player.Player) *session.Session

//go:linkname session_writePacket github.com/df-mc/dragonfly/server/session.(*Session).writePacket
func session_writePacket(*session.Session, packet.Packet)
