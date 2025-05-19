package instructions

import (
	"github.com/df-mc/dragonfly/server/session"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"time"
)

const (
	EasingTypeLinear = iota
	EasingTypeSpring
	EasingTypeInQuad
	EasingTypeOutQuad
	EasingTypeInOutQuad
	EasingTypeInCubic
	EasingTypeOutCubic
	EasingTypeInOutCubic
	EasingTypeInQuart
	EasingTypeOutQuart
	EasingTypeInOutQuart
	EasingTypeInQuint
	EasingTypeOutQuint
	EasingTypeInOutQuint
	EasingTypeInSine
	EasingTypeOutSine
	EasingTypeInOutSine
	EasingTypeInExpo
	EasingTypeOutExpo
	EasingTypeInOutExpo
	EasingTypeInCirc
	EasingTypeOutCirc
	EasingTypeInOutCirc
	EasingTypeInBounce
	EasingTypeOutBounce
	EasingTypeInOutBounce
	EasingTypeInBack
	EasingTypeOutBack
	EasingTypeInOutBack
	EasingTypeInElastic
	EasingTypeOutElastic
	EasingTypeInOutElastic
)

type CameraEase struct {
	// Type is the type of easing function used. This is one of the constants above.
	Type uint8
	// Duration is the time that the easing function should take.
	Duration time.Duration
}

type SetCameraInstruction struct {
	// Preset is one of the pre-defined presets that you can find in the presets.go file.
	Preset protocol.CameraPreset
	// Ease represents the easing function that is used by the instruction.
	Ease protocol.Optional[CameraEase]
	// Position represents the position of the camera.
	Position protocol.Optional[mgl32.Vec3]
	// Rotation represents the rotation of the camera.
	Rotation protocol.Optional[mgl32.Vec2]
	// Facing is a vector that the camera will always face towards during the duration of the instruction.
	Facing protocol.Optional[mgl32.Vec3]
	// ViewOffset is an offset based on a pivot point to the player, causing the camera to be shifted in a
	// certain direction.
	ViewOffset protocol.Optional[mgl32.Vec2]
	// EntityOffset is an offset from the entity that the camera should be rendered at.
	EntityOffset protocol.Optional[mgl32.Vec3]
}

func (i SetCameraInstruction) Packet(*session.Session) packet.Packet {
	var ease protocol.Optional[protocol.CameraEase]
	if e, ok := i.Ease.Value(); ok {
		ease = protocol.Option(protocol.CameraEase{
			Type:     e.Type,
			Duration: float32(e.Duration.Seconds()),
		})
	}

	return &packet.CameraInstruction{
		Set: protocol.Option(protocol.CameraInstructionSet{
			Preset:       indexes[i.Preset],
			Ease:         ease,
			Position:     i.Position,
			Rotation:     i.Rotation,
			Facing:       i.Facing,
			ViewOffset:   i.ViewOffset,
			EntityOffset: i.EntityOffset,
		}),
	}
}
